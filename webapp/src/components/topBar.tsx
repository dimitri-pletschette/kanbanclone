// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react'

import './topBar.scss'
import {FormattedMessage} from 'react-intl'

import HelpIcon from '../widgets/icons/help'
import {Utils} from '../utils'
import {Constants} from '../constants'

const TopBar = React.memo((): JSX.Element => {
    if (Utils.isFocalboardPlugin()) {
        const feedbackUrl = 'https://www.focalboard.com/fwlink/feedback-boards.html?v=' + Constants.versionString
        return (
            <div
                className='TopBar'
            >
                <a
                    className='link'
                    href={feedbackUrl}
                >
                    <FormattedMessage
                        id='TopBar.give-feedback'
                        defaultMessage='Give Feedback'
                    />
                </a>
                <div className='versionFrame'>
                    <div className='version'>
                        {`v${Constants.versionString}`}
                    </div>
                    <div className='versionBadge'>{'BETA'}</div>
                </div>
            </div>
        )
    }

    const focalboardFeedbackUrl = 'https://www.focalboard.com/fwlink/feedback-focalboard.html?v=' + Constants.versionString
    return (
        <div
            className='TopBar'
        >
            <a
                className='link'
                href={focalboardFeedbackUrl}
            >
                <FormattedMessage
                    id='TopBar.give-feedback'
                    defaultMessage='Give Feedback'
                />
            </a>
            <a
                href='https://www.focalboard.com/guide/user?utm_source=webapp'
                target='_blank'
                rel='noreferrer'
            >
                <HelpIcon/>
            </a>
        </div>
    )
})

export default TopBar
