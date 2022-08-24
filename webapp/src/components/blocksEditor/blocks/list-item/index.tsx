import React, {useRef, useEffect} from 'react'
import {BlockInputProps, ContentType} from '../types'

import './list-item.scss'

const ListItem: ContentType = {
    name: 'list-item',
    displayName: 'List item',
    slashCommand: '/list-item',
    prefix: '* ',
    nextType: 'list-item',
    render: (value: string) => <ul><li>{value}</li></ul>,
    runSlashCommand: (): void => {},
    Input: (props: BlockInputProps) => {
        const ref = useRef<HTMLInputElement|null>(null)
        useEffect(() => {
            ref.current?.focus()
        }, [])
        return (
            <ul>
                <li>
                    <input
                        ref={ref}
                        className='ListItem'
                        onChange={(e) => props.onChange(e.currentTarget.value)}
                        onKeyDown={(e) => {
                            if (props.value === '' && e.key === "Backspace") {
                                props.onCancel()
                            }
                            if (e.key === "Enter") {
                                props.onSave(props.value)
                            }
                        }}
                        value={props.value}
                    />
                </li>
            </ul>
        )
    }
}

ListItem.runSlashCommand = (changeType: (contentType: ContentType) => void, changeValue: (value: string) => void, ...args: string[]): void => {
    changeType(ListItem)
    changeValue(args.join(' '))
}

export default ListItem
