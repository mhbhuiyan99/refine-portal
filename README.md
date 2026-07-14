# Property Refine Page

A property search and refine page built with **Go (Beego)** and **Vanilla JavaScript**. The project integrates OwnerDirect APIs to search locations, retrieve properties, display property details, and provide sorting and filtering functionality.

## Features

- Search properties by location
- Property listing with server-side rendering
- Dynamic property tiles
- Property details integration
- Sort properties
  - Most Popular
  - Price (Low → High)
  - Price (High → Low)
  - Highest Rating
  - Lowest Rating
- Refine filters
  - Date
  - Price Range
  - Guests
  - Pet Friendly (UI)
  - Eco Friendly (UI)
- Currency conversion based on country
- Responsive UI

## Tech Stack

- Go
- Beego
- HTML
- CSS
- Vanilla JavaScript

## APIs Used

- Location API
- Properties API
- Property Details API

## Project Structure

```
controllers/
services/
routers/
views/
static/
    css/
    js/
    images/
```

## Run the Project

```bash
go mod tidy
```

```bash
bee run
```

Open your browser:

```
http://localhost:8080/refine?search=dhaka,Bangladesh
```
