// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {useCallback, useEffect} from 'react'
import {generatePath, useRouteMatch, useHistory} from 'react-router-dom'
import {FormattedMessage} from 'react-intl'

import {getCurrentBoard} from '../store/boards'
import {getCurrentViewCardsSortedFilteredAndGrouped} from '../store/cards'
import {getView, getCurrentBoardViews, getCurrentViewGroupBy, getCurrentView} from '../store/views'
import {useAppSelector, useAppDispatch} from '../store/hooks'

import {Utils} from '../utils'

import {getClientConfig, setClientConfig} from '../store/clientConfig'

import wsClient, {WSClient} from '../wsclient'
import {ClientConfig} from '../config/clientConfig'

import CenterPanel from './centerPanel'
import EmptyCenterPanel from './emptyCenterPanel'

import Sidebar from './sidebar/sidebar'
import './workspace.scss'

type Props = {
    readonly: boolean
}

function CenterContent(props: Props) {
    const match = useRouteMatch<{boardId: string, viewId: string, cardId?: string}>()
    const board = useAppSelector(getCurrentBoard)
    const cards = useAppSelector(getCurrentViewCardsSortedFilteredAndGrouped)
    const activeView = useAppSelector(getView(match.params.viewId))
    const views = useAppSelector(getCurrentBoardViews)
    const groupByProperty = useAppSelector(getCurrentViewGroupBy)
    const clientConfig = useAppSelector(getClientConfig)
    const history = useHistory()
    const dispatch = useAppDispatch()

    const showCard = useCallback((cardId?: string) => {
        const params = {...match.params, cardId}
        const newPath = generatePath(match.path, params)
        history.push(newPath)
    }, [match, history])

    useEffect(() => {
        const onConfigChangeHandler = (_: WSClient, config: ClientConfig) => {
            Utils.log('dispatch')
            dispatch(setClientConfig(config))
        }
        wsClient.addOnConfigChange(onConfigChangeHandler)

        // wsClient.addOnReconnect(onReconnect)
        return () => {
            wsClient.removeOnConfigChange(onConfigChangeHandler)

            // wsClient.removeOnReconnect(onReconnect)
        }
    }, [])

    if (board && activeView) {
        Utils.log('ENABLE' + (clientConfig?.enablePublicSharedBoards || 'NULL'))
        let property = groupByProperty
        if ((!property || property.type !== 'select') && activeView.fields.viewType === 'board') {
            property = board?.fields.cardProperties.find((o) => o.type === 'select')
        }
        return (
            <CenterPanel
                readonly={props.readonly}
                board={board}
                cards={cards}
                shownCardId={match.params.cardId}
                showCard={showCard}
                activeView={activeView}
                groupByProperty={property}
                views={views}
                showShared={clientConfig?.enablePublicSharedBoards || false}
            />
        )
    }

    return (
        <EmptyCenterPanel/>
    )
}

const Workspace = React.memo((props: Props) => {
    const board = useAppSelector(getCurrentBoard)
    const view = useAppSelector(getCurrentView)

    return (
        <div className='Workspace'>
            {!props.readonly &&
                <Sidebar
                    activeBoardId={board?.id}
                    activeViewId={view?.id}
                />
            }
            <div className='mainFrame'>
                {(board?.fields.isTemplate) &&
                <div className='banner'>
                    <FormattedMessage
                        id='Workspace.editing-board-template'
                        defaultMessage="You're editing a board template"
                    />
                </div>}
                <CenterContent readonly={props.readonly}/>
            </div>
        </div>
    )
})

export default Workspace
