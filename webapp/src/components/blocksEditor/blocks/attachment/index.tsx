import React, {useRef, useEffect, useState} from 'react'
import {BlockInputProps, ContentType} from '../types'
import octoClient from '../../../../octoClient'
import {useAppSelector} from '../../../../store/hooks'
import {getCurrentBoardId} from '../../../../store/boards'

import './attachment.scss'

type FileInfo = {
    file: string|File
    filename: string
}

const Attachment: ContentType<FileInfo> = {
    name: 'attachment',
    displayName: 'Attachment',
    slashCommand: '/attachment',
    prefix: '',
    runSlashCommand: (): void => {},
    editable: false,
    Display: (props: BlockInputProps<FileInfo>) => {
        const [fileDataUrl, setFileDataUrl] = useState<string|null>(null)
        const boardId = useAppSelector(getCurrentBoardId)

        useEffect(() => {
            if (!fileDataUrl) {
                const loadFile = async () => {
                    if (props.value && props.value.file && typeof props.value.file === 'string') {
                        const fileURL = await octoClient.getFileAsDataUrl(boardId, props.value.file)
                        setFileDataUrl(fileURL.url || '')
                    }
                }
                loadFile()
            }
        }, [props.value, props.value.file, boardId])

        return (
            <div className='AttachmentView'>
                <a
                    href={fileDataUrl || '#'}
                    onClick={(e) => e.stopPropagation()}
                    download={props.value.filename}
                >
                    📎 {props.value.filename}
                </a>
            </div>
        )
    },
    Input: (props: BlockInputProps<FileInfo>) => {
        const ref = useRef<HTMLInputElement|null>(null)
        useEffect(() => {
            ref.current?.click()
        }, [])

        return (
            <input
                ref={ref}
                className='Attachment'
                type='file'
                onChange={(e) => {
                    const files = e.currentTarget?.files
                    if (files) {
                        for (let i = 0; i < files.length; i++) {
                            const file = files.item(i)
                            if (file) {
                                const file = (e.currentTarget?.files || [])[0]
                                props.onSave({file: file, filename: file.name})
                            }
                        }
                    }
                }}
            />
        )
    }
}

Attachment.runSlashCommand = (changeType: (contentType: ContentType<FileInfo>) => void, changeValue: (value: FileInfo) => void): void => {
    changeType(Attachment)
    changeValue({} as any)
}

export default Attachment
