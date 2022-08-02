// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {useState} from 'react'
import {useIntl} from 'react-intl'

import {IPropertyTemplate} from '../../blocks/board'
import {FilterClause} from '../../blocks/filterClause'
import {createFilterGroup} from '../../blocks/filterGroup'
import {BoardView} from '../../blocks/boardView'
import mutator from '../../mutator'
import {Utils} from '../../utils'
import Button from '../../widgets/buttons/button'
import Menu from '../../widgets/menu'
import Editable from '../../widgets/editable'
import MenuWrapper from '../../widgets/menuWrapper'
import {PropertyType} from '../../properties/types'

import './filterValue.scss'

type Props = {
    view: BoardView
    filter: FilterClause
    template: IPropertyTemplate
    propertyType: PropertyType
}

const filterValue = (props: Props): JSX.Element|null => {
    const {filter, template, view, propertyType} = props
    const [value, setValue] = useState(filter.values.length > 0 ? filter.values[0] : '')
    console.log(filter)
    console.log(value)
    const intl = useIntl()

    if (propertyType.filterValueType === 'none') {
        return null
    }

    if (propertyType.filterValueType === 'boolean') {
        return null
    }

    if (propertyType.filterValueType === 'options' && filter.condition !== 'includes' && filter.condition !== 'notIncludes') {
        return null
    }

    if (propertyType.filterValueType === 'text') {
        return (
            <Editable
                onChange={setValue}
                value={value}
                placeholderText={intl.formatMessage({id: 'FilterByText.placeholder', defaultMessage: 'filter text'})}
                onSave={() => {
                    const filterIndex = view.fields.filter.filters.indexOf(filter)
                    Utils.assert(filterIndex >= 0, "Can't find filter")

                    const filterGroup = createFilterGroup(view.fields.filter)
                    const newFilter = filterGroup.filters[filterIndex] as FilterClause
                    Utils.assert(newFilter, `No filter at index ${filterIndex}`)

                    newFilter.values = [value]
                    mutator.changeViewFilter(view.boardId, view.id, view.fields.filter, filterGroup)
                }}
            />
        )
    }

    let displayValue: string
    if (filter.values.length > 0) {
        displayValue = filter.values.map((id) => {
            const option = template.options.find((o) => o.id === id)
            return option?.value || '(Unknown)'
        }).join(', ')
    } else {
        displayValue = '(empty)'
    }

    return (
        <MenuWrapper className='filterValue'>
            <Button>{displayValue}</Button>
            <Menu>
                {template.options.map((o) => (
                    <Menu.Switch
                        key={o.id}
                        id={o.id}
                        name={o.value}
                        isOn={filter.values.includes(o.id)}
                        onClick={(optionId) => {
                            const filterIndex = view.fields.filter.filters.indexOf(filter)
                            Utils.assert(filterIndex >= 0, "Can't find filter")

                            const filterGroup = createFilterGroup(view.fields.filter)
                            const newFilter = filterGroup.filters[filterIndex] as FilterClause
                            Utils.assert(newFilter, `No filter at index ${filterIndex}`)
                            if (filter.values.includes(o.id)) {
                                newFilter.values = newFilter.values.filter((id) => id !== optionId)
                                mutator.changeViewFilter(view.boardId, view.id, view.fields.filter, filterGroup)
                            } else {
                                newFilter.values.push(optionId)
                                mutator.changeViewFilter(view.boardId, view.id, view.fields.filter, filterGroup)
                            }
                        }}
                    />
                ))}
            </Menu>
        </MenuWrapper>
    )
}

export default filterValue
