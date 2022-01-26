// Copyright 2021-2022 the Kubeapps contributors.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	chartmodels "github.com/kubeapps/kubeapps/pkg/chart/models"
	dbutils "github.com/kubeapps/kubeapps/pkg/dbutils"
)

type AssetManager interface {
	Init() error
	Close() error
	GetChart(namespace, chartID string) (chartmodels.Chart, error)
	GetChartVersion(namespace, chartID, version string) (chartmodels.Chart, error)
	GetChartFiles(namespace, filesID string) (chartmodels.ChartFiles, error)
	GetPaginatedChartListWithFilters(cq ChartQuery, pageNumber, pageSize int) ([]*chartmodels.Chart, int, error)
	GetAllChartCategories(cq ChartQuery) ([]*chartmodels.ChartCategory, error)
}

// ChartQuery is a container for passing the supported query parameters for generating the WHERE query
type ChartQuery struct {
	Namespace   string
	ChartName   string
	Version     string
	AppVersion  string
	SearchQuery string
	Repos       []string
	Categories  []string
}

func NewManager(databaseType string, config dbutils.Config, globalReposNamespace string) (AssetManager, error) {
	return NewPGManager(config, globalReposNamespace)
}
