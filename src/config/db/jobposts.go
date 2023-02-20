package db

import "github.com/KadirbekSharau/apprentice-backend/src/models"

var jobTypes = []models.JobPostType{
	{
		Name: "Jobs",
	},
	{
		Name: "Internships",
	},
}

var locations = []models.Location{
	{
		Country: "Kazakhstan",
		City:    "Astana",
	},
	{
		Country: "Kazakhstan",
		City:    "Almaty",
	},
}

var companies = []models.Company{
	{
		UserID:             1,
		CompanyName:        "Google",
		CompanyDescription: "Google LLC is an American multinational technology company focusing on search engine technology, online advertising, cloud computing, computer software, quantum computing, e-commerce, artificial intelligence,[9] and consumer electronics.",
		CompanyWebsiteUrl:  "careers.google.com",
	},
	{
		UserID:           1,
		CompanyName:      "Tinkoff Bank",
	},
	{
		UserID:           1,
		CompanyName:      "One Technologies",
	},
	{
		UserID:           1,
		CompanyName:      "KazDream Technologies",
	},
	{
		UserID:           1,
		CompanyName:      "Akvelon",
	},
	{
		UserID:           1,
		CompanyName:      "Kolesa Group",
	},
	{
		UserID:           1,
		CompanyName:      "Kaspi.kz",
	},
}

var jobPosts = []models.JobPost{
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
	{
		UserID:        1,
		JobPostTypeID: 2,
		Name:          "Software Engineering Internship",
		LocationID:    1,
		CompanyID:     2,
		IsRemote:      false,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
}
