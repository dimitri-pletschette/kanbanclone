// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {createSlice, PayloadAction} from '@reduxjs/toolkit'

import {BoardView} from '../blocks/boardView'

import {initialLoad} from './initialLoad'

import {RootState} from './index'

const viewsSlice = createSlice({
    name: 'views',
    initialState: {views: []} as {views: BoardView[]},
    reducers: {
        updateViews: (state, action: PayloadAction<BoardView[]>) => {
            const updatedViewIds = action.payload.map((o: BoardView) => o.id)
            const newViews = state.views.filter((o: BoardView) => !updatedViewIds.includes(o.id)).map((o) => new BoardView(o))
            const updatedAndNotDeletedViews = action.payload.filter((o: BoardView) => o.deleteAt === 0)
            newViews.push(...updatedAndNotDeletedViews)
            state.views = newViews.sort((a, b) => a.title.localeCompare(b.title)).map((v) => new BoardView(v))
        },
    },
    extraReducers: (builder) => {
        builder.addCase(initialLoad.fulfilled, (state, action) => {
            state.views = action.payload.blocks.filter((o) => o.type === 'view').sort((a, b) => a.title.localeCompare(b.title)).map((o) => new BoardView(o))
        })
    },
})

export const {updateViews} = viewsSlice.actions
export const {reducer} = viewsSlice

export function getViews(state: RootState): BoardView[] {
    return state.views.views
}

export function getView(viewId: string): (state: RootState) => BoardView|null {
    return (state: RootState): BoardView|null => {
        return state.views.views.find((v) => v.id === viewId) || null
    }
}

export function getBoardViews(boardId: string): (state: RootState) => BoardView[] {
    return (state: RootState): BoardView[] => {
        return state.views.views.filter((v) => v.parentId === boardId)
    }
}
