package model

type Space struct {
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
	Free  int     `json:"free"`
	Total int     `json:"total"`
}

type SpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

//func FloatToSpacesResponse(spaces []float64) SpacesResponse {
//	resp := SpacesResponse{}
//	total := 0
//
//	for i := 0; i < len(spaces); i += 2 {
//		space := Space{
//			Lat:  spaces[i],
//			Long: spaces[i+1],
//		}
//		resp.Spaces = append(resp.Spaces, space)
//		total += 1
//	}
//	resp.Totals = total
//
//	return resp
//}
