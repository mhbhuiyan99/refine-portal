# Refine Portal

> A property discovery portal built with **Beego** demonstrating API integration, dynamic filtering, sorting, and responsive UI components.

**Project Type:** Beego Web Application  

---

## Overview

Refine Portal is a full-stack property discovery application that integrates with external property APIs to showcase rental properties across multiple locations. The project demonstrates best practices in **MVC architecture**, **API integration**, **server-side rendering (SSR)**, and **client-side rendering (CSR)** using Go, Beego, and modern JavaScript.

---

## Core Features

### Task 1: Refine Page (Dynamic Search & Filter)

A modern property search interface with real-time filtering and sorting capabilities.

**Features:**
- Destination Autocomplete - Search properties by location in real-time
- Property Grid - Responsive 4-column desktop, 2-column tablet, 1-column mobile layout
- Dynamic Filtering - Filter by property type, price range, amenities, and more
- Smart Sorting - Sort by relevance, price, rating, and popularity
- Property Cards - Display images, ratings, amenities, pricing, and partner info
- Currency Formatting - Dynamic price conversion based on user locale
- Partner Integration - Direct booking links via Booking.com, Vrbo, Expedia
- Breadcrumb Navigation - Easy location hierarchy navigation

**URL Structure:**
```
GET /refine?search=Dhaka,Bangladesh&order=1
```

---

### Task 2: Category Page (Location-wise Browsing)

Server-side rendered dynamic location pages for browsing properties by geography.

**Features:**
- Server-Side Rendering (SSR) - Fast initial page load with SEO optimization
- Dynamic Location Pages - Hierarchical location support (country, state, city)
- Property Sections - Organized property listings by category
- Category Filtering - Property type and amenity filtering
- Hero Section - Location-specific imagery and metadata
- Responsive Design - Seamless experience across all devices
- Breadcrumb Navigation - Location hierarchy visualization

**URL Structure:**
```
GET /all/:location
GET /all/usa
GET /all/usa/texas
GET /all/bangladesh/dhaka-division/dhaka
```

---

### Task 3: Request Layer Refactor

A dedicated `requests` layer was introduced to cleanly separate external API communication from business logic and routing.

**Architecture changes:**
- `controllers/` now handle request routing, parameter validation, and response rendering only
- `services/` now orchestrate business logic and delegate external API reads to the request layer
- `requests/` now contains all HTTP API logic: URL construction, headers, query parameters, request execution, response decoding, and error handling

**Why this matters:**
- Reduces duplicated HTTP client code across the app
- Keeps service boundary clear and controllers simple
- Enables reusable request helpers for all external endpoints
- Makes new API endpoints easier to add and maintain

**Backend data flow:**
1. Controller receives an incoming HTTP request
2. Controller extracts and validates route/query parameters
3. Controller calls a service function
4. Service calls the request layer to fetch external data
5. Request layer builds the HTTP request, executes it, and decodes the response
6. Service returns typed models back to the controller
7. Controller renders JSON or template output

**What changed:**
- `requests/client.go` now contains shared HTTP helpers: `BuildURL`, `NewGETRequest`, and `DoRequest`
- `requests/location_request.go`, `requests/property_list_request.go`, `requests/property_request.go`, and `requests/category_request.go` now use shared URL construction and response parsing
- `services/` files are now compact and focus on orchestration rather than request plumbing
- `controllers/` now translate incoming parameters and render responses only

---

## Configuration

The API base URL is configurable and should be set in `conf/app.conf`. To keep documentation generic, use placeholder host values in examples.

- `base_url` - Base API host for external requests
- `image_base_url` - Base URL used for images

**Example config values:**
```ini
base_url = https://api.example.com
image_base_url = https://images.example.com/640x287/
```

---

## API Integration

Refine Portal integrates with three main external property APIs, all configured through the `base_url` setting:

### Location API
**Purpose:** Destination search & autocomplete  
**Endpoint:** `GET /api/location/v1`  
**Parameters:**
- `keyword` - Search term (e.g., "dhaka, Bangladesh")
- `isLocationEntity` - Boolean to filter location entities

**Example:**
```
<BASE_URL>/api/location/v1
  ?keyword=dhaka,Bangladesh
  &isLocationEntity=true
```

---

### Property List API
**Purpose:** Retrieve property IDs and metadata for a location  
**Endpoint:** `GET /api/properties/category/v1`  
**Parameters:**
- `category` - Location category path (e.g., "bangladesh/dhaka-division/dhaka/973")
- `order` - Sort order (1 = relevance)
- `limit` - Number of properties to return
- `page` - Pagination page number
- `locations` - Location codes (e.g., "BD")
- `device` - Device type (desktop/mobile)

**Example:**
```
<BASE_URL>/api/properties/category/v1
  ?order=1
  &category=bangladesh/dhaka-division/dhaka/973
  &limit=192
  &items=1
  &locations=BD
  &device=desktop
  &page=1
```

---

### Property Details API
**Purpose:** Fetch complete property information (images, prices, amenities, ratings)  
**Endpoint:** `GET /api/property/bookmark/v1`  
**Parameters:**
- `propertyIdList` - Comma-separated property IDs from Property List API

**Returns:**
- Property images
- Price information
- Ratings & reviews
- Amenities & features
- Partner booking URLs
- Property feed ID

**Example:**
```
<BASE_URL>/api/property/bookmark/v1
  ?propertyIdList=prop123,prop456,prop789
```

---

### Category API (Location-wise Details)
**Purpose:** Retrieve category metadata, hero section, and aggregated property data  
**Endpoint:** `GET /api/v1/category/details`  
**Parameters:**
- `category` - Category path (e.g., "usa:texas")
- `aggsAvgPrice` - Include average price aggregation
- `aggsAvgRating` - Include average rating aggregation
- `aggsAvgRoomSize` - Include average room size aggregation
- `aggsCategory` - Include category aggregation
- `device` - Device type
- `locations` - Location codes

**Returns:**
- Category metadata & descriptions
- Hero section information
- Property sections & aggregations
- Statistics (avg price, rating, room size)

**Example:**
```
<BASE_URL>/api/v1/category/details/usa:texas
  ?aggsAvgPrice=1
  &aggsAvgRating=1
  &aggsAvgRoomSize=1
  &aggsCategory=1
  &device=desktop
  &items=1
  &locations=US
  &sections=1
```

---

## Architecture & Data Flow

### Refine Page (Client-Side Rendering)
```
┌─────────────┐
│   Browser   │ User enters search term
└──────┬──────┘
       │
       ▼
┌──────────────────┐
│ Location API     │ Autocomplete suggestions
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Property List    │ Get property IDs for location
│ API              │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Property Details │ Fetch full property info
│ API              │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ JSON Response    │
└──────┬───────────┘
       │
       ▼
┌──────────────────────┐
│ Client-Side Renderer │ JavaScript renders to DOM
└──────┬───────────────┘
       │
       ▼
┌──────────────────┐
│ Property Cards   │ Display in responsive grid
└──────────────────┘
```

### Category Page (Server-Side Rendering)
```
┌─────────────┐
│   Browser   │ Navigate to /all/usa/texas
└──────┬──────┘
       │
       ▼
┌──────────────────┐
│ Category         │
│ Controller       │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Category API     │ Fetch location data & properties
│                  │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Data Processing  │ Format & organize data
│ (Go)             │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Template Context │ Store in Beego context
│                  │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ Beego Template   │ Server-side rendering (TPL)
│ Engine           │
└──────┬───────────┘
       │
       ▼
┌──────────────────┐
│ HTML Page        │ Fully rendered HTML sent to browser
└──────────────────┘
```

---

## Key Components

### Property Card Component
**Used in:** Refine page, Category page  
**Shared across:** Both Task 1 & Task 2

**Displays:**
- Property image with fallback
- Property type badge
- Star rating with count
- Bed, bath, guest count
- Price per night (with currency)
- Location/neighborhood
- Amenities (with icons)
- Partner logo & "View Deal" button
- Direct booking link to partner

**Variants:**
- Desktop: Full details visible
- Mobile: Condensed layout

---

### Filter & Sort Components

**Filtering:**
- Property type filtering
- Price range slider
- Amenities multi-select
- Guest count selection
- Bedroom/bathroom filters
- Date range picker

**Sorting Options:**
- Relevance (default)
- Price (low to high)
- Price (high to low)
- Rating (highest first)
- Newest properties

---

## Partner Integration

Dynamic partner linking based on property feed ID:

| Feed ID | Partner | Logo | Link |
|---------|---------|------|------|
| 11 | Booking.com | Yes | booking.com/property |
| 12 | Vrbo | Yes | vrbo.com/property |
| 24 | Expedia | Yes | expedia.com/property |

Both partner logo and **"View Deal"** button use the partner's booking URL.

---

## Responsive Design

**Layout Breakpoints:**

| Device | Columns | Width |
|--------|---------|-------|
| Desktop | 4 | 1200px+ |
| Tablet | 2 | 768px - 1199px |
| Mobile | 1 | < 768px |

**Technologies:** CSS Grid, Flexbox, Mobile-first media queries

---

## Breadcrumb Navigation

Enables hierarchical navigation across locations.

**Refine Page Example:**
```
Home > Bangladesh > Dhaka Division > Dhaka
```

**Category Page Example:**
```
Home > USA > Texas > Austin
```

---

## Technologies & Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| **Backend** | Go 1.25 | Programming language |
| **Framework** | Beego v2 | Web framework & MVC |
| **Frontend** | HTML5 | Markup & structure |
| **Styling** | CSS3 | Responsive design |
| **JavaScript** | Vanilla JS | Client-side logic & DOM manipulation |
| **API** | REST | Integration with external property APIs |
| **Rendering** | SSR (Beego) + CSR (JS) | Dual rendering strategies |

---

## Client-Side JavaScript Features

- Destination autocomplete with API calls
- Dynamic property grid rendering
- Property type tab switching
- Real-time search filtering
- Multi-criteria sorting
- Filter modal interactions
- Currency conversion & formatting
- Amenities icon rendering
- Date/guest selection modals
- State management for filters


---

## Project Structure

```
refine-portal/
├── main.go                      # Application entry point
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies lock file
├── README.md                    # This file
│
├── conf/
│   ├── app.conf                 # Main application configuration
│   └── app.conf.example         # Configuration template
│
├── controllers/
│   ├── refine.go                # Refine page controller (Task 1)
│   ├── category.go              # Category page controller (Task 2)
│   ├── location_api.go          # Location API handler
│   └── property_api.go          # Property API handler
│
├── models/
│   ├── category.go              # Category data models
│   ├── location.go              # Location data models
│   ├── property.go              # Property data models
│   └── property_details.go      # Property details data models
│
├── services/                    # Business logic & orchestration
│   ├── location_service.go      # Location API service
│   ├── property_service.go      # Property List API service
│   ├── property_details_service.go  # Property Details API service
│   ├── category_service.go      # Category API service
│   └── helper.go                # Utility helper functions
│
├── requests/                    # External API request layer
│   ├── client.go                # Shared HTTP client and request helpers
│   ├── location_request.go      # Location API request logic
│   ├── property_list_request.go # Property List API request logic
│   ├── property_request.go      # Property Details API request logic
│   └── category_request.go      # Category API request logic
│
├── routers/
│   └── router.go                # Route definitions & configuration
│
├── views/                       # Beego template files
│   ├── refine.tpl               # Refine page template (Task 1)
│   ├── category.tpl             # Category page template (Task 2)
│   ├── components/
│   │   └── property_card.tpl    # Shared property card component
│   └── layouts/
│       ├── header.tpl           # Header layout
│       └── footer.tpl           # Footer layout
│
└── static/                      # Frontend assets
    ├── css/
    │   ├── refine.css           # Refine page styles
    │   ├── filter.css           # Filter modal styles
    │   ├── category.css         # Category page styles
    │   └── components/
    │       └── property_card.css # Property card component styles
    ├── js/
    │   ├── refine.js            # Refine page logic
    │   ├── category.js          # Category page logic
    │   ├── filter.js            # Filter functionality
    │   ├── filter_modal.js      # Filter modal interactions
    │   ├── filter_apply.js      # Filter application logic
    │   ├── filter_state.js      # Filter state management
    │   ├── sort.js              # Sorting functionality
    │   ├── api.js               # API client utilities
    │   ├── renderer.js          # DOM rendering utilities
    │   ├── date_modal.js        # Date picker modal
    │   ├── guest_model.js       # Guest selection modal
    │   ├── components/
    │   │   ├── property_card.js # Property card rendering
    │   │   ├── header.js        # Header component
    │   │   ├── navbar.js        # Navigation bar
    │   │   ├── breadcrumb.js    # Breadcrumb navigation
    │   │   └── sort.js          # Sort selector
    │   └── utils/
    │       ├── amenity_icons.js # Amenity icon mapping
    │       ├── currency.js      # Currency formatting
    │       └── partner_logo.js  # Partner logo utilities
    └── images/
        └── amenities/           # Amenity icons
```

---

## Getting Started

### Prerequisites

- **Go** 1.25 or higher
- **Beego v2** framework
- **Git** for version control
- **Internet connection** (for API access)

### Installation

#### 1. Clone the Repository
```bash
git clone https://github.com/mhbhuiyan99/refine-portal.git
cd refine-portal
```

#### 2. Install Dependencies
```bash
go mod tidy
```

#### 3. Configure Environment

Copy the example configuration file and update with your credentials:

```bash
cp conf/app.conf.example conf/app.conf
```

Edit `conf/app.conf`:

```ini
appname = refine-portal
httpport = 8080
runmode = dev

# Base URLs for API and images
base_url = https://api.example.com
image_base_url = https://images.example.com/640x287/

api_key = <YOUR_API_KEY>
basic_auth_username = <USERNAME>
basic_auth_password = <PASSWORD>
```


#### 4. Run the Application

**Using Go (direct):**
```bash
go run main.go
```

**Using Bee (Beego CLI, optional):**
```bash
bee run
```

The application will start on `http://localhost:8080`

---

## Available Routes

### Refine Page (Task 1)
**Route:** `GET /refine`

**Query Parameters:**
- `search` - Search location (e.g., "Dhaka, Bangladesh")
- `order` - Sort order (1 = relevance, 2 = price low-high, etc.)

**Examples:**
```
http://localhost:8080/refine
http://localhost:8080/refine?search=Dhaka,Bangladesh&order=1
http://localhost:8080/refine?search=New%20York&order=2
```

---

### Category Page (Task 2)
**Route:** `GET /all/:location`

**Examples:**
```
http://localhost:8080/all/usa
http://localhost:8080/all/usa/texas
http://localhost:8080/all/bangladesh
http://localhost:8080/all/bangladesh/dhaka-division
http://localhost:8080/all/bangladesh/dhaka-division/dhaka
```

---

## Development Workflow

### Code Organization

**Controllers** - Receive HTTP requests and render responses
**Services** - Orchestrate business flow and delegate API calls to the request layer
**Requests** - Handle all external HTTP API communication, headers, URL building, and response decoding
**Models** - Define data structures
**Views** - Render HTML templates
**Static Assets** - CSS, JavaScript, images

### Adding a New Feature

1. **Create Model** - Define data structure in `models/`
2. **Create Service** - Implement business logic in `services/`
3. **Create Request Logic** - Add external API request logic in `requests/`
4. **Create Controller** - Handle HTTP requests in `controllers/`
5. **Add Route** - Register route in `routers/router.go`
6. **Create Template** - Design UI in `views/`
7. **Add Styles & Scripts** - CSS in `static/css/`, JS in `static/js/`

### API Integration Pattern

```go
// Request layer - handles external API HTTP calls
func GetLocationRequest(keyword string) (*models.LocationResponse, error) {
    // Build URL
    // Create HTTP request
    // Execute request
    // Decode response
}

// Service layer - orchestrates business logic
func GetLocation(keyword string) (*models.LocationResponse, error) {
    return requests.GetLocationRequest(keyword)
}

// Controller layer - handles HTTP
func (c *LocationAPIController) Get() {
    location, err := services.GetLocation(keyword)
    c.Data["json"] = location
    c.ServeJSON()
}
```

---

## Testing

Run your changes locally:

```bash
# Terminal 1: Start server
go run main.go

# Terminal 2: Test endpoints
curl http://localhost:8080/refine?search=Dhaka
curl http://localhost:8080/all/usa
```



## Performance Optimization

### Implemented:
- Skeleton loading cards with shimmer effect for property listings
- Header integration for improved loading UX
- Infinite scrolling and load more functionality for property results
- Optimized client-side rendering for property cards
- Lazy loading for images
- CSS Grid for efficient layout
- Minified static assets
- Efficient API client reuse
- Concurrent property details fetching using goroutines and wait groups

### Future Improvements:
- Response caching layer (Redis)
- Image CDN optimization
- Server-side request deduplication
- GraphQL API integration
- WebSocket for real-time updates

---

## Documentation

### File Purpose Reference

| File | Purpose |
|------|---------|
| [main.go](main.go) | Application bootstrap & initialization |
| [routers/router.go](routers/router.go) | Route registration & middleware |
| [controllers/refine.go](controllers/refine.go) | Refine page request handler |
| [controllers/category.go](controllers/category.go) | Category page request handler |
| [requests/client.go](requests/client.go) | Centralized HTTP client and request helpers for all API calls |
| [requests/location_request.go](requests/location_request.go) | Location API request logic |
| [requests/property_list_request.go](requests/property_list_request.go) | Property List API request logic |
| [requests/property_request.go](requests/property_request.go) | Property Details API request logic |
| [requests/category_request.go](requests/category_request.go) | Category API request logic |
| [static/js/refine.js](static/js/refine.js) | Refine page JavaScript logic |
| [static/js/category.js](static/js/category.js) | Category page JavaScript logic |
| [views/refine.tpl](views/refine.tpl) | Refine page template |
| [views/category.tpl](views/category.tpl) | Category page template |

---

## Implementation Checklist

### Task 1: Refine Page
- Create Beego project structure
- Implement `/refine` route and controller
- Create `refine.tpl` template
- Integrate Location API for autocomplete
- Integrate Property List API
- Integrate Property Details API
- Implement property grid rendering
- Implement filter functionality
- Implement sort functionality
- Create property card component
- Add responsive design
- Add partner logo integration
- Add breadcrumb navigation

### Task 2: Category Page
- Create `/all/:location` routes
- Create category page controller
- Create `category.tpl` template
- Integrate Category API
- Implement server-side rendering
- Add hero section
- Add property sections
- Implement filter functionality
- Create responsive layout
- Reuse property card component

### Task 3: Request Layer Refactor
- Add `requests/` package for all external API calls
- Move HTTP request logic out of controllers and services
- Create shared request helpers for headers, URL building, and response parsing
- Keep services small and focused on orchestration
- Keep controllers limited to HTTP handling and response formatting
- Update docs to describe the new request-layer architecture


---

## Notes

- The application uses **server-side rendering** for the Category page.
- The Refine page renders property cards dynamically using JavaScript after fetching data from backend APIs.
- Backend controllers act as a proxy to external APIs, handling authentication and response processing.
- Property images, partner information, breadcrumbs, and location details are normalized on the backend before being rendered.