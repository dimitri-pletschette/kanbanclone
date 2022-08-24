import React, {useRef, useEffect} from 'react'
import {BlockInputProps, ContentType} from '../types'

import './divider.scss'

const Divider: ContentType = {
    name: 'divider',
    displayName: 'Divider',
    slashCommand: '/divider',
    prefix: '--- ',
    render: () => <hr className='Divider'/>,
    runSlashCommand: (): void => {},
    Input: (props: BlockInputProps) => {
        useEffect(() => {
            props.onSave(props.value)
        }, [])
        return null
    }
}

Divider.runSlashCommand = (changeType: (contentType: ContentType) => void, changeValue: (value: string) => void, ...args: string[]): void => {
    changeType(Divider)
    changeValue(args.join(' '))
}

export default Divider
