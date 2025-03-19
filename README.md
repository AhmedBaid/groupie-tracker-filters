# Groupie Tracker filters

Groupie Tracker filter is a Go-based web application that interacts with a RESTful API to fetch and display data about musical artists, their concert locations, dates, and relationships. This project aims to provide an intuitive and structured way to visualize this information.

## 📌 Project Overview

Groupie Tracker is a web application that fetches and presents data about music artists, their concerts, and related details. Users can explore artist information, past and upcoming concert locations, and their relationships with other events.

## Objectives

The application connects to an API containing four key parts:

- **Artists**: Information about bands and artists, including their names, images, formation year, first album date, and members.
- **Locations**: Locations of their past and/or upcoming concerts.
- **Dates**: Dates of their past and/or upcoming concerts.
- **Relations**: Links between artists, dates, and locations.

The main goal is to build a website that presents this information using various UI components, such as cards, lists, and tables.

## Features

✅ Fetch and display artist details dynamically.
✅ Show concert dates and locations in a structured manner.
✅ Provide a user-friendly interface.
✅ Implement efficient client-server communication.
✅ Handle error pages gracefully.

## 🛠 Technology Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML/CSS for UI
- **API**: RESTful API for data retrieval

---

## 💾 Project Structure

The project follows a modular structure for better organization and maintainability:

```
GROUPIE-TRACKER/
│── cmd/
│   └── main.go              # Entry point of the application
│
├── handler/                 # Handles HTTP requests
│   ├── Detail_Card_Func.go  # Handles artist details page
│   ├── Groupie_Func.go      # Handles main page requests
│   ├── Style_Func.go        # Manages styles for the website
│
├── helpers/                 # Helper functions for data fetching and rendering
│   ├── fetching_By_Id.go      # Fetches data by artist ID
│   ├── page_Deleted.go       # Handles deleted pages
│   ├── render_Template.go   # Utility for rendering templates
│
├── routes/
│   ├── routes.go            # Handles routing for the application
│
├── static/                  # Static files (CSS, Images, etc.)
│   ├── images/              # Stores static images
│   ├── card_Detail.css      # Styling for artist details page
│   ├── index.css            # Styling for homepage
│   ├── status_Page.css      # Styling for status/error pages
│
├── template/                # HTML templates
│   ├── details_Card.html     # Template for artist details
│   ├── index.html           # Homepage template
│   ├── status_Page.html      # Error/status page template
│
├── tools/                   # Utility functions and data structures
│   ├── Tools.go             # Contains various helper functions
│
├── go.mod                   # Go module file
├── Dockerfile               # Docker setup for deployment
└── README.md                # Project documentation
```

---

## 🚀 Installation & Setup

### Prerequisites
- Install [Go](https://go.dev/)

### Steps
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/groupie-tracker.git
   cd groupie-tracker
   ```
2. Initialize Go modules:
   ```sh
   go mod tidy
   ```
3. Run the project:
   ```sh
   go run cmd/main.go
   ```
4. Open your browser and visit:
   ```
   http://localhost:8080
   ```

---

## 📺 API Source
This project fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/).

## 📄 License
This project is open-source and available under the [MIT License](LICENSE).

