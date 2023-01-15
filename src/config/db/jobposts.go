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

var businessStreams = []models.BusinessStream{
	{
		BusinessStreamName: "IT",
	},
}

var companies = []models.Company{
	{
		UserID:             1,
		BusinessStreamID:   1,
		CompanyName:        "Google",
		CompanyDescription: "Google LLC is an American multinational technology company focusing on search engine technology, online advertising, cloud computing, computer software, quantum computing, e-commerce, artificial intelligence,[9] and consumer electronics.",
		CompanyWebsiteUrl:  "careers.google.com",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "Tinkoff Bank",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "One Technologies",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "KazDream Technologies",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "Akvelon",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "Kolesa Group",
	},
	{
		UserID:           1,
		BusinessStreamID: 1,
		CompanyName:      "Kaspi.kz",
	},
}

var skillSets = []models.SkillSet{
	{
		Name: "Analytics",
	},
	{
		Name: "C++",
	},
	{
		Name: "Golang",
	},
	{
		Name: "Frontend",
	},
	{
		Name: "Backend",
	},
	{
		Name: "IOS",
	},
	{
		Name: "Android",
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
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
		IsPaid:        true,
		Description:   "Analytics, QA-Engineering, SRE, Java, Frontend , Scala, .NET, Go, ML-engineer, Python, C++, Android, IOS Development",
		Link:          "https://github.com/danabeknar/kazakhstan-it-internships",
		IsActive:      true,
	},
}
