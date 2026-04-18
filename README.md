<div align="center">

# 📈 Stock-Ed Trading Company

### *Learn. Trade. Succeed.*

**A full-stack trading education website — Go HTTP server + HTML/CSS/JS frontend**

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![Google Fonts](https://img.shields.io/badge/Google_Fonts-4285F4?style=for-the-badge&logo=google&logoColor=white)

</div>

---

## 🌐 Overview

**Stock-Ed** is a stock market trading education platform built for the Indian market. The project is structured as a **layer-based Go web server** that serves premium dark-themed HTML pages for the Stock-Ed Trading Company — covering courses, company vision, team, roadmap, and student enrollment.

> *"From basics to smart level, we cover everything you need for profitable trading."*

---

## 📁 Project Structure

```
github.com/abhinandpn/Website
├── cmd/
│   └── server/
│       └── main.go               # Entry point — wires routes & starts HTTP server
│
├── internal/
│   ├── config/
│   │   └── config.go             # Reads PORT env variable, defaults to 8080
│   ├── handler/
│   │   └── handler.go            # Page handlers & contact form POST endpoint
│   └── middleware/
│       └── middleware.go         # Logger, Panic Recovery, CORS middleware
│
├── web/
│   ├── static/
│   │   ├── styles.css            # Global shared stylesheet
│   │   └── assets/
│   │       ├── logo.png          # Brand logo (also used as favicon)
│   │       ├── chart.png         # Trading chart dashboard image
│   │       ├── team_meeting.png  # Team banner image
│   │       └── avatar.png        # Placeholder student avatar
│   └── templates/
│       ├── index.html            # Landing page
│       ├── about.html            # About / company page
│       └── form.html             # Registration & enrollment form
│
├── go.mod                        # Go module: github.com/abhinandpn/Website
├── .gitignore
└── README.md
```

---

## 🖥️ Pages

### `/` — Landing Page (`index.html`)
- **Hero Section** — Bold headline, welcome block with trading philosophy
- **Value Prop Card** — Glowing card with chart image and key offering
- **About Preview** — Team image with animated green glow overlay
- **Features Section** — Pill badges + detailed feature items (Assignments, Personal Attention)
- **Testimonials** — Infinite dual-direction marquee scroll
- **Contact Section** — Embedded quick-enquiry form
- **Footer** — CTA, navigation links, branding

### `/about` — About Page (`about.html`)
- **Hero** — Company tagline and philosophy
- **Our Vision** — Animated glowing vision card
- **Mentorship Methods** — 01 / 02 / 03 methodology layout
- **Team Grid** — 9 team member cards
- **Roadmap** — "Road to 1 Crore" animated SVG path
- **Segments** — 3 glowing pill modules (Regular, Trading, Live Sessions)

### `/form` — Registration Form (`form.html`)
- **Multi-field form** — Name, Place, Contact, Email, Current Status, Course
- **Trading Background** — Radio card selector (Beginner / Intermediate / Advanced)
- **Privacy Consent** — Checkbox agreement
- **Submit** — Posts to `POST /contact`

---

## ⚙️ Backend — Go Server

### Architecture

```
Request → Middleware Chain → Router (ServeMux) → Handler → Template → Response
```

### Layer Breakdown

| Layer | Package | Responsibility |
|-------|---------|----------------|
| **Entry** | `cmd/server` | Boot the server, wire routes and middleware |
| **Config** | `internal/config` | Read `PORT` env var, return defaults |
| **Handler** | `internal/handler` | Serve HTML templates, handle form POSTs |
| **Middleware** | `internal/middleware` | Request Logger, Panic Recovery, CORS |
| **Frontend** | `web/` | HTML templates & static assets |

### API Routes

| Method | Route | Description |
|--------|-------|-------------|
| `GET` | `/` | Landing page |
| `GET` | `/about` | About page |
| `GET` | `/form` | Registration form page |
| `GET` | `/static/*` | Static files (CSS, images) |
| `POST` | `/contact` | Contact form submission handler |

### Middleware Stack

```
Recovery  →  CORS  →  Logger  →  Handler
```
- **Recovery** — Catches panics, logs stack trace, returns `500`
- **CORS** — Adds permissive headers for local development
- **Logger** — Logs `[METHOD] /path — duration` for every request

---

## 🎨 Design System

| Token | Value | Usage |
|-------|-------|-------|
| `--bg-primary` | `#030805` | Page backgrounds |
| `--accent` | `#00ff00` | Neon green highlights & glows |
| `--font-heading` | `Playfair Display` | All headings (h1–h4) |
| `--font-body` | `Inter` | Body text, labels, buttons |
| `--transition` | `0.4s cubic-bezier(...)` | All hover & state transitions |
| `--text-body` | `#e0e0e0` | Paragraph text colour |

### Key Visual Features
- 🌑 Deep dark `#030805` background with neon `#00ff00` accent
- ✨ Smooth cinematic scroll-reveal on all major sections (`IntersectionObserver`)
- 🌟 Gradient border glow effects on key cards (`::before` pseudo-element)
- 🎞️ Infinite dual-direction marquee for testimonials
- 💚 Animated `text-shadow` glow on headings
- 🔤 `Playfair Display` serif headings + `Inter` body — premium typography

---

## 🚀 Getting Started

### Prerequisites
- [Go 1.22+](https://golang.org/dl/)

### Run Locally

```bash
# Clone the repository
git clone https://github.com/abhinandpn/Website.git
cd Website

# Run the Go server
go run cmd/server/main.go
```

Open your browser at **http://localhost:8080**

### Custom Port

```bash
PORT=3000 go run cmd/server/main.go
```

### Build Binary

```bash
go build -o bin/server cmd/server/main.go
./bin/server
```

---

## 🔗 Navigation Map

| Page | URL | Links To |
|------|-----|----------|
| Home | `http://localhost:8080/` | `/about`, `/form`, `#contact` |
| About | `http://localhost:8080/about` | `/` (Home), `#contact` |
| Form | `http://localhost:8080/form` | `/` (Home) |

---

## 📞 Contact Information

- 📱 **Phone:** +91 90377 13791
- 📧 **Email:** stockedhelp@gmail.com

---

## 👨‍💻 Author

**Designed & Developed by** `abhinand_pn`

**Module:** `github.com/abhinandpn/Website`

---

<div align="center">

© STOCK-ED. 2024 · All rights reserved.

</div>
