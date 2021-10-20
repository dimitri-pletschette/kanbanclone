// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React from 'react'

import {render, screen} from '@testing-library/react'

import {createMemoryHistory} from 'history'

import {Router} from 'react-router-dom'

import userEvent from '@testing-library/user-event'

import {wrapIntl} from '../../testUtils'

import WelcomePage from './welcomePage'

describe('pages/welcome', () => {
    const history = createMemoryHistory()
    test('Welcome Page shows Explore Page', () => {
        const {container} = render(wrapIntl(
            <Router history={history}>
                <WelcomePage/>
            </Router>,
        ))
        expect(screen.getByText('Explore')).toBeDefined()
        expect(container).toMatchSnapshot()
    })

    test('Welcome Page shows Explore Page And Then Proceeds after Clicking Explore', () => {
        history.replace = jest.fn()
        render(wrapIntl(
            <Router history={history}>
                <WelcomePage/>
            </Router>,
        ))
        const exploreButton = screen.getByText('Explore')
        expect(exploreButton).toBeDefined()
        userEvent.click(exploreButton)
        expect(history.replace).toBeCalledWith('/dashboard')
    })

    test('Welcome Page does not render explore page the second time we visit it', () => {
        render(wrapIntl(
            <Router history={history}>
                <WelcomePage/>
            </Router>,
        ))
        expect(history.replace).toBeCalledWith('/dashboard')
    })

    test('Welcome Page redirects us when we have a r query parameter', () => {
        history.replace = jest.fn()
        history.location.search = 'r=123'
        render(wrapIntl(
            <Router history={history}>
                <WelcomePage/>
            </Router>,
        ))
        expect(history.replace).toBeCalledWith('123')
    })
})
