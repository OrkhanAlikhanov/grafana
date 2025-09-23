package loki

import (
	"time"

	common "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1"
	"github.com/grafana/grafana/pkg/tsdb/loki/kinds/dataquery"
)

type QueryType = dataquery.LokiQueryType
type SupportingQueryType = dataquery.SupportingQueryType
type Direction = dataquery.LokiQueryDirection

const (
	QueryTypeRange   = dataquery.LokiQueryTypeRange
	QueryTypeInstant = dataquery.LokiQueryTypeInstant
)

const (
	SupportingQueryLogsVolume                         = dataquery.SupportingQueryTypeLogsVolume
	SupportingQueryLogsSample                         = dataquery.SupportingQueryTypeLogsSample
	SupportingQueryDataSample                         = dataquery.SupportingQueryTypeDataSample
	SupportingQueryInfiniteScroll                     = dataquery.SupportingQueryTypeInfiniteScroll
	SupportingQueryNone           SupportingQueryType = "none"
)

const (
	DirectionBackward = dataquery.LokiQueryDirectionBackward
	DirectionForward  = dataquery.LokiQueryDirectionForward
)

type lokiQuery struct {
	Expr                string
	QueryType           QueryType
	Direction           Direction
	Step                time.Duration
	MaxLines            int
	LegendFormat        string
	Start               time.Time
	End                 time.Time
	RefID               string
	SupportingQueryType SupportingQueryType
	Scopes              []common.ScopeFilter
}
