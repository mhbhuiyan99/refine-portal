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

The final architecture uses a layered design that clearly separates responsibilities across the application.

**Architecture:**

- **Controllers**
  - Handle routing, request validation, HTTP context, and template/JSON responses.
  - No longer communicate directly with external APIs.
- **Services**
  - Contain business logic and orchestrate application flow.
  - Coordinate calls to the request layer.
  - Handle execution timing and service-level logging.
  - Keep controllers lightweight and focused on HTTP handling.
- **Requests**
  - Responsible for all external API communication.
  - Build request URLs, create HTTP requests, parse responses, and return typed models.
- **Shared HTTP Client**
  - `requests/client.go` centralizes request creation, common headers, authentication, query parameters, response parsing, error handling, and logging.

**Benefits:**

- Clear separation of concerns.
- Business logic moved from controllers into services.
- Reusable HTTP request utilities across all API layers.
- Easier maintenance and better readability.
- Easier to add new APIs.
- Consistent logging and error handling.

**Backend data flow:**

1. Controller receives an incoming HTTP request.
2. Controller validates parameters and request context.
3. Controller calls a service function.
4. Service executes business logic and calls the request layer.
5. Request layer builds and executes the HTTP request, parses the response, and returns typed models.
6. Service returns structured data back to the controller.
7. Controller renders JSON or templates.

**What changed:**

- `requests/client.go` now provides shared HTTP helpers for every external endpoint.
- `requests/` handles all API-specific request building and response parsing.
- `services/` orchestrate requests and business logic, keeping controllers small.
- `controllers/` now only handle HTTP flow and rendering.
- `requests/property_image_request.go` supports the Property Images API.

---

### Task 4: Property Image Slider

Property cards now include a shared image slider component used by both the Refine Page and Category Page.

**Implementation details:**

- Property images are fetched **on demand**.
- Initially only the feature image is displayed.
- Additional images are requested only when the user clicks the next arrow.
- A three-dot loading indicator appears while the image request is in progress.
- Images are cached after the first request to avoid repeated API calls for the same property.
- The same slider implementation is shared by both Refine and Category pages.
- Navigation supports previous/next buttons and dot indicators.
- The slider remains responsive and works across different screen sizes.

**Why this matters:**

- Reduces initial page weight by delaying image loading.
- Avoids duplicate API requests for the same property.
- Maintains a responsive, consistent UI across pages.
- Gives users smooth image browsing inside each property card.

---

### Task 5: Currency Switcher

A currency dropdown in the shared header lets users switch prices across the entire app, with the selection persisted across the Refine and Category pages.

**Implementation details:**

- A currency dropdown lives in the shared header, so it's available on both the Refine Page and Category Page.
- Duplicate currencies are collapsed in the dropdown — countries that share a currency (e.g. several EU countries using EUR) appear as a single option.
- Selecting a currency instantly re-renders all visible property prices, including the price range filter.
- Prices are converted from a fixed USD base price using each currency's exchange rate, then formatted with the correct symbol.
- The selected currency is saved so it persists across a page reload and when navigating between the Refine and Category pages.
- On page load, the previously saved currency (or a sensible default) is applied automatically before prices are shown.
- The conversion and formatting logic is centralized in one place and reused by both pages, avoiding duplicated pricing code.

**Why this matters:**

- One shared currency system instead of separate logic per page.
- Prices, symbols, and the price filter always stay in sync with the selected currency.
- Persisting the selection avoids forcing users to reselect their currency on every page or reload.

---

### Task 6: Category & Sub-Category Routing

The Category Page supports two separate route types:

* **Category routes** for normal location-based pages.
* **Sub-category routes** for predefined sub-category pages.

The two request types are handled by separate controllers.

**Route structure:**

```text
Category:
GET /all/*

Examples:
GET /all/bangladesh
GET /all/bangladesh/dhaka-division
GET /all/bangladesh/dhaka-division/dhaka


Sub-category:
GET /all/*/<sub-category>

Examples:
GET /all/bangladesh/pet-friendly
GET /all/bangladesh/pools
GET /all/bangladesh/luxury
GET /all/bangladesh/beach
```

**Routing behavior:**

Predefined sub-category routes are registered before the general category route.

```text
Browser
   │
   │ GET /all/bangladesh/pet-friendly
   ▼
Beego Router
   │
   ├── Matches predefined sub-category route
   │
   ▼
SubCategoryController
```

For a normal location URL:

```text
Browser
   │
   │ GET /all/bangladesh/dhaka
   ▼
Beego Router
   │
   ├── No predefined sub-category route matches
   │
   ▼
CategoryController
```

This keeps category and sub-category responsibilities separate.

**Sub-category route registration:**

```go
for _, slug := range services.SubCategorySlugs() {
    web.Router(
        "/all/*/"+slug,
        &controllers.SubCategoryController{},
    )
}

web.Router("/all/*", &controllers.CategoryController{})
```

The predefined sub-category slugs are obtained from the sub-category registry. This prevents the controller from having to determine whether the final URL segment represents a location or a sub-category.

**Sub-category request flow:**

```text
Browser
   │
   │ GET /all/bangladesh/pet-friendly
   ▼
Beego Router
   │
   ▼
SubCategoryController
   │
   ├── Extract URL path
   ├── Extract last segment: "pet-friendly"
   ├── Resolve sub-category from registry
   ├── Extract location: "bangladesh"
   ├── Resolve country code
   │
   ▼
Sub-category Registry
   │
   └── pet-friendly
          ↓
       petFriendly
          ↓
       amenities=11
   │
   ▼
CategoryService
   │
   ▼
CategoryRequest
   │
   ├── Build category slug: "bangladesh"
   ├── Add locations=BD
   ├── Add amenities=11
   │
   ▼
Category API
   │
   ▼
CategoryResponse
   │
   ▼
SubCategoryController
   │
   ▼
category.tpl
```

**Normal category request flow:**

```text
Browser
   │
   │ GET /all/bangladesh/dhaka
   ▼
Beego Router
   │
   ▼
CategoryController
   │
   ├── Extract location URL
   ├── Convert "/" to ":"
   ├── Resolve country code
   │
   ▼
CategoryService
   │
   ▼
CategoryRequest
   │
   ▼
Category API
   │
   ▼
CategoryResponse
   │
   ▼
CategoryController
   │
   ▼
category.tpl
```

**Sub-category mapping:**

The sub-category registry remains the single source of truth for supported sub-categories, URL aliases, canonical keys, and API parameters.

```text
URL slug              Canonical key       API parameter
--------------------------------------------------------
pet-friendly          petFriendly         amenities=11
pools                 pools               amenities=12
luxury                luxuryRental        order=3
beach                 beachRental         amenities=18-19
family                familyRental        amenities=5
```

The controller does not hard-code these API filter values.

**Architecture:**

```text
                    Beego Router
                         │
              ┌──────────┴──────────┐
              │                     │
       /all/*/<slug>              /all/*
              │                     │
              ▼                     ▼
   SubCategoryController     CategoryController
              │                     │
              └──────────┬──────────┘
                         │
                         ▼
               CategoryService
                         │
                         ▼
               CategoryRequest
                         │
                         ▼
                  Category API
```

This separation makes the routing behavior explicit and keeps the responsibilities of category and sub-category controllers independent.


### Additional Features

The final implementation also includes improved interactive filtering and pricing.

- Dynamic currency conversion.
- Guest filtering.
- Price filtering.
- Amenity filtering.
- Dynamic total price calculation.
- Property cards display the total price for a **7-night stay** by default.
- When the user selects check-in and check-out dates, the total price automatically updates to reflect the **actual number of selected nights**.
- The "/night" price always represents the nightly rate, while the line below dynamically shows the total stay cost.

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

**Purpose:** Destination search 
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

### Property Images API

**Purpose:** Fetch additional images for a single property to power the image slider
**Endpoint:** `GET /api/property/images/v1`  
**Parameters:**

- `propertyId` - Unique property identifier

**Returns:**

- A list of image file names for the requested property

**Example:**

```
<BASE_URL>/api/property/images/v1
  ?propertyId=12345
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
- `sections` - Include category property sections
- `amenities` - Filter properties by amenity ID
- `order` - Apply category sorting/filtering
- `pt` - Filter by property type
- `pax` - Filter by guest capacity
- `isWinter` / `isSummer` - Seasonal category filters
- `isShortTermStays` - Short-term stay filter
- `isBusinessTravel` - Business travel filter
- `ecoFriendly` - Sustainable/eco-friendly filter
- `sqs` - Additional category scope/filter

For sub-category URLs, these parameters are resolved from the sub-category registry instead of being hard-coded in the controller.

**Example:**

```text
URL:
GET /all/bangladesh/pet-friendly

Resolved location:
bangladesh

Resolved sub-category:
petFriendly

Category API request:
GET /api/v1/category/details/bangladesh
    ?amenities=11
    &locations=BD
    &sections=1
    ...
```


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
│ Location API     │ Get location info
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

| Feed ID | Partner     | Logo | Link                 |
| ------- | ----------- | ---- | -------------------- |
| 11      | Booking.com | Yes  | booking.com/property |
| 12      | Vrbo        | Yes  | vrbo.com/property    |
| 24      | Expedia     | Yes  | expedia.com/property |

Both partner logo and **"View Deal"** button use the partner's booking URL.

---

## Responsive Design

**Layout Breakpoints:**

| Device  | Columns | Width          |
| ------- | ------- | -------------- |
| Desktop | 4       | 1200px+        |
| Tablet  | 2       | 768px - 1199px |
| Mobile  | 1       | < 768px        |

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

| Layer          | Technology             | Purpose                                 |
| -------------- | ---------------------- | --------------------------------------- |
| **Backend**    | Go 1.25                | Programming language                    |
| **Framework**  | Beego v2               | Web framework & MVC                     |
| **Frontend**   | HTML5                  | Markup & structure                      |
| **Styling**    | CSS3                   | Responsive design                       |
| **JavaScript** | Vanilla JS             | Client-side logic & DOM manipulation    |
| **API**        | REST                   | Integration with external property APIs |
| **Rendering**  | SSR (Beego) + CSR (JS) | Dual rendering strategies               |

---

## Client-Side JavaScript Features

- Dynamic property grid rendering
- Property type tab switching
- Real-time search filtering
- Multi-criteria sorting
- Filter modal interactions
- Currency conversion, formatting & cross-page persistence (Task 5)
- Amenities icon rendering
- Date/guest selection modals
- State management for filters

---

## Project Structure

```
refine-portal/
├── main.go                      # Beego app bootstrap
├── go.mod                       # Go module definition
├── go.sum                       # Go dependency lock file
├── README.md                    # Project overview and usage docs
├── TEST.md                      # Testing notes and task checklist
│
├── conf/
│   ├── app.conf                 # Local runtime configuration
│   └── app.conf.example        # Sample configuration template
│
├── controllers/
│   ├── category.go              # Category page controller
│   ├── error.go                 # Shared Beego error rendering helpers
│   ├── location_api.go          # Location API controller
│   ├── property_api.go          # Property list/details API controller
│   ├── property_image_api.go    # Property images API controller
│   ├── refine.go                # Refine page controller
│   └── sub_category.go          # Sub-category route controller
│
├── models/
│   ├── category.go              # Category response models
│   ├── location.go              # Location response models
│   ├── property.go              # Property list/detail models
│   ├── property_details.go      # Property details models
│   └── property_image.go        # Property image models
│
├── requests/
│   ├── category_request.go      # Category API HTTP request logic
│   ├── client.go                # Shared HTTP client helpers
│   ├── config.go                # API config helpers
│   ├── location_request.go      # Location API request logic
│   ├── property_image_request.go # Property image request logic
│   ├── property_list_request.go # Property list request logic
│   ├── property_request.go      # Property details request logic
│   └── ...
│
├── services/
│   ├── category_service.go      # Category orchestration service
│   ├── helper.go                # Shared service helpers
│   ├── location_service.go      # Location service
│   ├── property_details_service.go # Property details service
│   ├── property_image_service.go # Property image service
│   ├── property_service.go      # Property list service
│   └── subcategory_registry.go  # Sub-category alias and parameter mapping
│
├── routers/
│   └── router.go                # Route registration for web pages and APIs
│
├── views/
│   ├── category.tpl             # Category page template
│   ├── refine.tpl              # Refine page template
│   ├── components/
│   │   └── property_card.tpl    # Shared property card component
│   ├── errors/
│   │   ├── 400.tpl             # Bad request error page
│   │   ├── 404.tpl             # Not found page
│   │   └── 500.tpl             # Internal server error page
│   └── layouts/
│       ├── footer.tpl          # Footer layout
│       └── header.tpl          # Header layout
│
├── static/
│   ├── css/
│   │   ├── category.css        # Category page styles
│   │   ├── components/
│   │   │   └── property_card.css # Shared card styling
│   │   ├── error.css           # Error page styles
│   │   ├── filter.css          # Filter modal styles
│   │   ├── refine.css          # Refine page styles
│   │   └── shimmer.css         # Loading shimmer effect
│   ├── images/
│   │   ├── amenities/          # Amenity icons and related assets
│   │   └── demo/               # Demo images and placeholders
│   └── js/
│       ├── api.js              # API helpers
│       ├── category.js         # Category page client logic
│       ├── components/
│       │   ├── breadcrumb.js   # Breadcrumb renderer
│       │   ├── header.js       # Shared header logic
│       │   ├── navbar.js       # Navbar/currency UI
│       │   ├── property_card.js # Property card rendering
│       │   └── sort.js         # Sort control logic
│       ├── date_modal.js       # Date picker modal logic
│       ├── filter.js           # Filter orchestration
│       ├── filter_apply.js     # Filter application logic
│       ├── filter_modal.js     # Filter modal behavior
│       ├── filter_state.js     # Saved filter state
│       ├── guest_model.js      # Guest count modal
│       ├── refine.js           # Refine page behavior
│       ├── refine_reload.js    # Refine refresh helper
│       ├── renderer.js         # DOM renderer utilities
│       └── utils/
│           ├── amenity_icons.js # Amenity icon mapping
│           ├── currency.js     # Currency formatting and conversion
│           ├── partner_logo.js # Logo selection helper
│           └── pricing.js      # Price display formatting
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

### Property Images API

**Route:** `GET /api/property/images/v1`

**Example:**

```
http://localhost:8080/api/property/images/v1?propertyId=12345
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

### Coverage

```bash
go tool cover -func=coverage.out
```

Current coverage output:

```text
refine-portal/controllers/category.go:27:               Get                             86.2%
refine-portal/controllers/category.go:130:              buildRefineURL                  100.0%
refine-portal/controllers/error.go:11:                  renderError                     100.0%
refine-portal/controllers/error.go:37:                  renderNotFound                  100.0%
refine-portal/controllers/error.go:41:                  renderBadRequest                100.0%
refine-portal/controllers/error.go:45:                  renderServerError               100.0%
refine-portal/controllers/location_api.go:20:           Get                             100.0%
refine-portal/controllers/property_api.go:24:           GetList                         100.0%
refine-portal/controllers/property_api.go:142:          GetDetails                      100.0%
refine-portal/controllers/property_image_api.go:20:     Get                             100.0%
refine-portal/controllers/refine.go:15:                 Get                             100.0%
refine-portal/controllers/sub_category.go:26:           Get                             95.2%
refine-portal/main.go:9:                                main                            0.0%
refine-portal/requests/category_request.go:23:          GetCategoryRequest              100.0%
refine-portal/requests/client.go:29:                    Error                           100.0%
refine-portal/requests/client.go:40:                    DoRequest                       100.0%
refine-portal/requests/client.go:92:                    BuildURL                        100.0%
refine-portal/requests/client.go:112:                   BuildImageURL                   100.0%
refine-portal/requests/client.go:131:                   NewGETRequest                   83.3%
refine-portal/requests/client.go:155:                   setDefaultHeaders               88.9%
refine-portal/requests/config.go:16:                    GetURLFromConfig                100.0%
refine-portal/requests/location_request.go:23:          GetLocationRequest              80.0%
refine-portal/requests/property_image_request.go:20:    GetPropertyImagesRequest        81.2%
refine-portal/requests/property_list_request.go:25:     GetPropertyListRequest          73.0%
refine-portal/requests/property_request.go:25:          GetPropertyDetailsRequest       83.3%
refine-portal/routers/router.go:10:                     init                            0.0%
refine-portal/services/category_service.go:21:          GetCategory                     95.0%
refine-portal/services/helper.go:4:                     chunkStrings                    100.0%
refine-portal/services/location_service.go:17:          GetLocation                     100.0%
refine-portal/services/property_details_service.go:26:  GetPropertyDetails              97.6%
refine-portal/services/property_image_service.go:13:    GetPropertyImages               0.0%
refine-portal/services/property_service.go:17:          GetProperties                   100.0%
refine-portal/services/subcategory_registry.go:166:     LookupSubCategory               0.0%
refine-portal/services/subcategory_registry.go:185:     SubCategorySlugs                0.0%

total:                                                  (statements)                    88.7%
```

This project is currently at 88.7% total statement coverage based on the generated `coverage.out` report.

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

| File                                                                     | Purpose                                                       |
| ------------------------------------------------------------------------ | ------------------------------------------------------------- |
| [main.go](main.go)                                                       | Application bootstrap & initialization                        |
| [routers/router.go](routers/router.go)                                   | Route registration & middleware                               |
| [controllers/refine.go](controllers/refine.go)                           | Refine page request handler                                   |
| [controllers/category.go](controllers/category.go)                       | Category page request handler                                 |
| [controllers/property_image_api.go](controllers/property_image_api.go)   | Property Images API handler                                    |
| [requests/client.go](requests/client.go)                                 | Centralized HTTP client and request helpers for all API calls |
| [requests/location_request.go](requests/location_request.go)             | Location API request logic                                    |
| [requests/property_list_request.go](requests/property_list_request.go)   | Property List API request logic                               |
| [requests/property_request.go](requests/property_request.go)             | Property Details API request logic                            |
| [requests/property_image_request.go](requests/property_image_request.go) | Property Images API request logic                             |
| [services/property_image_service.go](services/property_image_service.go)| Property Images service                                        |
| [requests/category_request.go](requests/category_request.go)             | Category API request logic                                    |
| [services/subcategory_registry.go](services/subcategory_registry.go)     | Centralized sub-category slug, alias, and API parameter mapping |
| [models/property_image.go](models/property_image.go)                     | Property Images data models                                    |
| [static/js/refine.js](static/js/refine.js)                               | Refine page JavaScript logic                                  |
| [static/js/category.js](static/js/category.js)                           | Category page JavaScript logic                                |
| [static/js/utils/currency.js](static/js/utils/currency.js)               | Currency list, rates & formatting helpers              |
| [static/js/utils/pricing.js](static/js/utils/pricing.js)                 | Applies the saved currency to rendered price tiles   |
| [static/js/components/navbar.js](static/js/components/navbar.js)         | Currency dropdown population & change handling       |
| [views/refine.tpl](views/refine.tpl)                                     | Refine page template                                          |
| [views/category.tpl](views/category.tpl)                                 | Category page template                                        |
| [views/layouts/header.tpl](views/layouts/header.tpl)                     | Shared navbar/header, including the currency dropdown         |

---

## Notes

- The application uses **server-side rendering** for the Category page.
- The Refine page renders property cards dynamically using JavaScript after fetching data from backend APIs.
- Backend controllers act as a proxy to external APIs, handling authentication and response processing.
- Property images, partner information, breadcrumbs, and location details are normalized on the backend before being rendered.