// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {useState, useRef, useEffect} from 'react'
import {FormattedMessage, useIntl} from 'react-intl'
import {useHotkeys} from 'react-hotkeys-hook'
import {debounce} from 'lodash'

import Button from '../../widgets/buttons/button'
import Editable from '../../widgets/editable'

import {useAppSelector, useAppDispatch} from '../../store/hooks'
import {getSearchText, setSearchText} from '../../store/searchText'

const ViewHeaderSearch = (): JSX.Element => {
    const searchText = useAppSelector<string>(getSearchText)
    const dispatch = useAppDispatch()
    const intl = useIntl()

    const searchFieldRef = useRef<{focus(selectAll?: boolean): void}>(null)
    const [isSearching, setIsSearching] = useState(Boolean(searchText))
    const [searchValue, setSearchValue] = useState(searchText)

    const dispatchSearchText = debounce((value: string) => {
        dispatch(setSearchText(value))
    }, 200)

    useEffect(() => {
        searchFieldRef.current?.focus()
    }, [isSearching])

    useHotkeys('ctrl+shift+f,cmd+shift+f', () => {
        setIsSearching(true)
        searchFieldRef.current?.focus(true)
    })

    if (isSearching) {
        return (
            <Editable
                ref={searchFieldRef}
                value={searchValue}
                placeholderText={intl.formatMessage({id: 'ViewHeader.search-text', defaultMessage: 'Search text'})}
                onChange={(value) => {
                    setSearchValue(value)
                    dispatchSearchText(value)
                }}
                onCancel={() => {
                    setSearchValue('')
                    setIsSearching(false)
                    dispatchSearchText('')
                }}
                onSave={() => {
                    if (searchValue === '') {
                        setIsSearching(false)
                    }
                    dispatchSearchText(searchValue)
                }}
            />
        )
    }
    return (
        <Button onClick={() => setIsSearching(true)}>
            <FormattedMessage
                id='ViewHeader.search'
                defaultMessage='Search'
            />
        </Button>
    )
}

export default ViewHeaderSearch
