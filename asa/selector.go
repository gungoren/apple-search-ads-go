/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asa

type ConditionOperator string

const (
	ConditionOperatorEquals      ConditionOperator = "EQUALS"
	ConditionOperatorGreaterThan ConditionOperator = "GREATER_THAN"
	ConditionOperatorLessThan    ConditionOperator = "LESS_THAN"
	ConditionOperatorIn          ConditionOperator = "IN"
	ConditionOperatorLike        ConditionOperator = "LIKE"
	ConditionOperatorStartsWith  ConditionOperator = "STARTSWITH"
	ConditionOperatorContains    ConditionOperator = "CONTAINS"
	ConditionOperatorEndsWith    ConditionOperator = "ENDSWITH"
	ConditionOperatorNotEqual    ConditionOperator = "NOT_EQUALS"
	ConditionOperatorIs          ConditionOperator = "IS"
	ConditionOperatorContainsAny ConditionOperator = "CONTAINS_ANY"
	ConditionOperatorContainsAll ConditionOperator = "CONTAINS_ALL"
)

type Selector struct {
	Conditions []*Condition `json:"conditions,omitempty"`
	Fields     []string     `json:"fields,omitempty"`
	OrderBy    []*Sorting   `json:"orderBy,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

type Condition struct {
	Field    string            `json:"field"`
	Operator ConditionOperator `json:"operator"`
	Values   []string          `json:"values"`
}

type SortOrder string

const (
	SortingOrderAscending  SortOrder = "ASCENDING"
	SortingOrderDescending SortOrder = "DESCENDING"
)

type Sorting struct {
	Field     string    `json:"field"`
	SortOrder SortOrder `json:"sortOrder"`
}

type Pagination struct {
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
}
