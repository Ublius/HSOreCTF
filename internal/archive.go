package internal

import (
	"net/http"
)

type Link struct {
	URL   string
	Title string
}

type WinningTeam struct {
	Place    string
	Name     string
	School   string
	Location string
}

type CompetitionResult struct {
	Name      string
	Shortname string
	Teams     []WinningTeam
}

type YearInfo struct {
	Year            int
	RecapParagraphs []string
	Links           []Link
	Results         []CompetitionResult
}

func (a *Application) GetArchiveTemplate(*http.Request) map[string]any {
	return map[string]any{
		"YearInfo": []YearInfo{
			{
				Year: 2025,
				RecapParagraphs: []string{
					"This was the second year the compettion was put on! Only an In-Person competition was held, with prizes being awarded to: first, second, and third place teams.",
					"The In-Person competition had 12 teams consisting of 48 total participants.",
				},
				Links: []Link{
					{"/static/solutions/2025-solutions.pdf", "Solution Sketch Slides"},
				},
				Results: []CompetitionResult{
					{
						Name:      "In-Person",
						Shortname: "InPerson",
						Teams: []WinningTeam{
							{"1st", "CCHS Team B", "Cherry Creek High School", "Greenwood Village, Colorado"},
							{"2nd", "<script>alert(1)</script>", "School Name", "Location"},
							{"3nd", "Good Question", "School Name", "Location"},
						},
					},
				},
			},
				{
					Year: 2024,
					RecapParagraphs: []string{
						"This was the first year the compettion was put held! Only an In-Person competition was held, with prizes being awarded to: first, and second place teams.",
						"The In-Person competition had 2 teams consisting of 8 total participants.",
					},
					Links: []Link{
						{"", ""},
					},
					Results: []CompetitionResult{
						{
							Name: "In-Person",
							Shortname: "In-Person",
							Teams: []WinningTeam{
								{"1st", "", "", ""},
								{"2nd", "", "", ""},
							},
						},
					},
				},
			// 	{
			// 		Year: 2021,
			// 		RecapParagraphs: []string{
			// 			"The 2021 competition was an all-remote competition featuring 55 teams from across the nation.",
			// 		},
			// 		Links: []Link{
			// 			{"https://sumnerevans.com/posts/school/2021-hspc/", "Competition Recap and Solution Sketches"},
			// 			{"https://mines21.kattis.com/problems", "Problems"},
			// 		},
			// 		Results: []CompetitionResult{
			// 			{
			// 				Teams: []WinningTeam{
			// 					{"1st", "River Hill HS Team 1", "River Hill High School", "Clarksville, Maryland"},
			// 					{"2nd", "PEN A Team", "PEN Academy", "Cresskill, New Jersey"},
			// 					{"3nd", "River Hill HS Team 2", "River Hill High School", "Clarksville, Maryland"},
			// 				},
			// 			},
			// 		},
			// 	},
			// 	{
			// 		Year: 2020,
			// 		RecapParagraphs: []string{
			// 			"Due to COVID, the 2020 competition was the first all-remote HSPC competition. The competition featured 30 teams.",
			// 		},
			// 		Links: []Link{
			// 			{"https://sumnerevans.com/posts/school/2020-hspc/", "Competition Recap and Solution Sketches"},
			// 			{"https://mines20.kattis.com/problems", "Problems"},
			// 		},
			// 		Results: []CompetitionResult{
			// 			{
			// 				Teams: []WinningTeam{
			// 					{"1st", "Installation Wizards", "STEM School Highlands Ranch", "Highlands Ranch, Colorado"},
			// 					{"2nd", "i", "STEM School Highlands Ranch", "Highlands Ranch, Colorado"},
			// 					{"3nd", "Sun Devils", "Kent Denver", "Denver, Colorado"},
			// 				},
			// 			},
			// 		},
			// 	},
			// 	{
			// 		Year: 2019,
			// 		RecapParagraphs: []string{
			// 			"The second ever CS@Mines High School Programming Competition featured 22 teams from all around Colorado and from as far as Steamboat Springs.",
			// 		},
			// 		Links: []Link{
			// 			{"https://sumnerevans.com/posts/school/2019-hspc/", "Competition Recap and Solution Sketches"},
			// 			{"https://mines19.kattis.com/problems", "Problems"},
			// 		},
			// 		Results: []CompetitionResult{
			// 			{
			// 				Teams: []WinningTeam{
			// 					{"1st", "STEM Team 1", "STEM School Highlands Ranch", "Highlands Ranch, Colorado"},
			// 					{"2nd", "IntrospectionExceptions", "Colorado Academy", "Lakewood, Colorado"},
			// 					{"3nd", "Team 2", "?", "?"},
			// 				},
			// 			},
			// 		},
			// 	},
			// 	{
			// 		Year: 2018,
			// 		RecapParagraphs: []string{
			// 			"The first ever CS@Mines High School Programming Competition featured 22 teams.",
			// 		},
			// 		Links: []Link{
			// 			{"https://mines18.kattis.com/problems", "Problems"},
			// 		},
			// 		Results: []CompetitionResult{
			// 			{
			// 				Teams: []WinningTeam{
			// 					{"1st", "The Crummies", "Warren Tech", "Arvada, Colorado"},
			// 					{"2nd", "The Bean Beans", "Colorado Academy", "Lakewood, Colorado"},
			// 					{"3nd", "Warriors", "Arapahoe High School", "Centennial, Colorado"},
			// 				},
			// 			},
			// 		},
			// 	},
		},
	}
}
