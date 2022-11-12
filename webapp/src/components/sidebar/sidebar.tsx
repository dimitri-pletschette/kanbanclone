// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {useEffect, useState, useContext, useCallback} from 'react'
import {FormattedMessage, useIntl} from 'react-intl'
import {useHistory, useRouteMatch} from 'react-router-dom'

import {getActiveThemeName, loadTheme} from '../../theme'
import mutator from '../../mutator'
import IconButton from '../../widgets/buttons/iconButton'
import HamburgerIcon from '../../widgets/icons/hamburger'
import HideSidebarIcon from '../../widgets/icons/hideSidebar'
import ShowSidebarIcon from '../../widgets/icons/showSidebar'
import isPagesContext from '../../isPages'
import {getCurrentBoardId} from '../../store/boards'
import {getMySortedBoards, getMySortedPageFolders} from '../../store/boards'
import {useAppDispatch, useAppSelector} from '../../store/hooks'
import {Utils} from '../../utils'
import {IUser} from '../../user'

import './sidebar.scss'

import {
    BoardCategoryWebsocketData,
    Category,
    CategoryBoards,
    fetchSidebarCategories,
    getSidebarCategories, updateBoardCategories,
    updateCategories,
} from '../../store/sidebar'

import BoardsSwitcher from '../boardsSwitcher/boardsSwitcher'

import wsClient, {WSClient} from '../../wsclient'

import {getCurrentTeam, Team} from '../../store/teams'

import {Constants} from '../../constants'

import {getMe} from '../../store/users'
import {getCurrentViewId} from '../../store/views'

import SidebarCategory from './sidebarCategory'
import SidebarSettingsMenu from './sidebarSettingsMenu'
import SidebarUserMenu from './sidebarUserMenu'

type Props = {
    activeBoardId?: string
    onBoardTemplateSelectorOpen: () => void
    onBoardTemplateSelectorClose?: () => void
}

function getWindowDimensions() {
    const {innerWidth: width, innerHeight: height} = window
    return {
        width,
        height,
    }
}

const Sidebar = (props: Props) => {
    const [isHidden, setHidden] = useState(false)
    const [userHidden, setUserHidden] = useState(false)
    const [windowDimensions, setWindowDimensions] = useState(getWindowDimensions())
    const currentBoardId = useAppSelector<string>(getCurrentBoardId) || null
    const currentTeam = useAppSelector<Team|null>(getCurrentTeam)
    const boards = useAppSelector(getMySortedBoards)
    const pages = useAppSelector(getMySortedPageFolders)
    const isPages = useContext(isPagesContext)
    const dispatch = useAppDispatch()
    const sidebarCategories = useAppSelector<CategoryBoards[]>(getSidebarCategories)
    const me = useAppSelector<IUser|null>(getMe)
    const activeViewID = useAppSelector(getCurrentViewId)
    const history = useHistory()
    const match = useRouteMatch<{boardId: string, viewId?: string}>()
    const intl = useIntl()

    useEffect(() => {
        wsClient.addOnChange((_: WSClient, categories: Category[]) => {
            dispatch(updateCategories(categories))
        }, 'category')

        wsClient.addOnChange((_: WSClient, blockCategories: BoardCategoryWebsocketData[]) => {
            dispatch(updateBoardCategories(blockCategories))
        }, 'blockCategories')
    }, [])

    const team = useAppSelector(getCurrentTeam)

    useEffect(() => {
        if (team) {
            dispatch(fetchSidebarCategories(team!.id))
        }
    }, [team?.id])

    useEffect(() => {
        loadTheme()
    }, [])

    useEffect(() => {
        function handleResize() {
            setWindowDimensions(getWindowDimensions())
        }

        window.addEventListener('resize', handleResize)
        return () => window.removeEventListener('resize', handleResize)
    }, [])

    useEffect(() => {
        hideSidebar()
    }, [windowDimensions])

    if (!boards) {
        return <div/>
    }

    const hideSidebar = () => {
        if (!userHidden) {
            if (windowDimensions.width < 768) {
                setHidden(true)
            } else {
                setHidden(false)
            }
        }
    }
    const showFolder = useCallback(async (boardId) => {
        Utils.showBoard(boardId, match, history)
    }, [match, history])

    if (!me) {
        return <div/>
    }

    if (isHidden) {
        return (
            <div className='Sidebar octo-sidebar hidden'>
                <div className='octo-sidebar-header show-button'>
                    <div className='hamburger-icon'>
                        <IconButton
                            icon={<HamburgerIcon/>}
                            onClick={() => {
                                setUserHidden(false)
                                setHidden(false)
                            }}
                        />
                    </div>
                    <div className='show-icon'>
                        <IconButton
                            icon={<ShowSidebarIcon/>}
                            onClick={() => {
                                setUserHidden(false)
                                setHidden(false)
                            }}
                        />
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div className='Sidebar octo-sidebar'>
            {!Utils.isFocalboardPlugin() &&
                <div className='octo-sidebar-header'>
                    <div className='heading'>
                        <SidebarUserMenu/>
                    </div>

                    <div className='octo-spacer'/>
                    <div className='sidebarSwitcher'>
                        <IconButton
                            onClick={() => {
                                setUserHidden(true)
                                setHidden(true)
                            }}
                            icon={<HideSidebarIcon/>}
                        />
                    </div>
                </div>}

            {team && team.id !== Constants.globalTeamId &&
                <div className='WorkspaceTitle'>
                    {Utils.isFocalboardPlugin() &&
                    <>
                        <div className='octo-spacer'/>
                        <div className='sidebarSwitcher'>
                            <IconButton
                                onClick={() => {
                                    setUserHidden(true)
                                    setHidden(true)
                                }}
                                icon={<HideSidebarIcon/>}
                            />
                        </div>
                    </>
                    }
                </div>
            }

            <BoardsSwitcher
                onBoardTemplateSelectorOpen={props.onBoardTemplateSelectorOpen}
                onFolderCreate={() => {
                    mutator.addEmptyFolder(currentTeam?.id || '', intl, showFolder, () => showFolder(currentBoardId))
                }}
                userIsGuest={me?.is_guest}
                isPages={isPages}
            />

            <div className='octo-sidebar-list'>
                {
                    sidebarCategories.map((category, index) => (
                        <SidebarCategory
                            hideSidebar={hideSidebar}
                            key={category.id}
                            activeBoardID={props.activeBoardId}
                            activeViewID={activeViewID}
                            categoryBoards={category}
                            boards={isPages ? pages : boards}
                            isPages={isPages}
                            allCategories={sidebarCategories}
                            index={index}
                            onBoardTemplateSelectorClose={props.onBoardTemplateSelectorClose}
                        />
                    ))
                }
            </div>

            <div className='octo-spacer'/>

            {
                (!Utils.isFocalboardPlugin()) &&
                <div
                    className='add-board'
                    onClick={props.onBoardTemplateSelectorOpen}
                >
                    <FormattedMessage
                        id='Sidebar.add-board'
                        defaultMessage='+ Add board'
                    />
                </div>
            }

            {!Utils.isFocalboardPlugin() &&
                <SidebarSettingsMenu activeTheme={getActiveThemeName()}/>}
        </div>
    )
}

export default React.memo(Sidebar)
