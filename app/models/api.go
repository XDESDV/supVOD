package models

// WSResponse est le format standardisé de la réponse.
//   - Meta         : *Entête pré-formaté d'une réponse retournant des données.*
//   - Data         : *Donnée ou liste de données retournée(s).*
type WSResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

// MetaResponse est une entête d'une réponse valide
//   - ObjectName  : *Information retourné au front lui permettant de savoir quel format il reçoit.*
//   - TotalCount  : *Nombre total d'enregistrement que la demande peut retourner.*
//   - Offset      : *Position de départ de la liste des enregistrements retournés au Front.*
//   - Count       : *Nombre d'enregistrement retourné au Front.*
type MetaResponse struct {
	ObjectName string `json:"object_name"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offSet"`
	Count      int    `json:"count"`
}

type PaginationParams struct {
	Page     int
	PageSize int
}

type SearchParams struct {
	Search string
	Kinds  []Kind
}
