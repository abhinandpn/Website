<div align="center">

# 📈 Stock-Ed Trading Company

### *Learn. Trade. Succeed.*

**A professional dark-themed trading education website built with HTML, CSS & Vanilla JavaScript**

![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![Google Fonts](https://img.shields.io/badge/Google_Fonts-4285F4?style=for-the-badge&logo=google&logoColor=white)

</div>

---

## 🌐 Overview

**Stock-Ed** is a trading education platform built for the Indian stock market. The website is designed for Stock-Ed Trading Company to showcase their courses, mentorship values, roadmap to success, and onboarding form — all wrapped in a premium cinematic dark-green UI.

> *"From basics to smart level, we cover everything you need for profitable trading."*

---

## 📁 Project Structure

```
Website/
├── index.html          # Landing page — Hero, Features, Testimonials, Contact form
├── about.html          # About page — Vision, Team, Roadmap, Segments
├── form.html           # Registration / Enrollment form page
├── styles.css          # Global shared stylesheet
├── assets/
│   ├── logo.png        # Stock-Ed brand logo (also used as favicon)
│   ├── chart.png       # Trading chart dashboard image
│   ├── team_meeting.png # Team photo / banner image
│   └── avatar.png      # Placeholder student profile avatar
└── README.md           # This file
```

---

## 🖥️ Pages

### `index.html` — Landing Page
The main entry point of the website. Includes:
- **Hero Section** — Bold headline, welcome text and value proposition
- **Value Prop Card** — Highlight card with chart image and key offering
- **About Preview** — Team image with green glow overlay
- **Features Section** — Pill badges + detailed feature items
- **Testimonials** — Infinite dual-direction marquee scroll
- **Contact Section** — Embedded enquiry form with name, email, phone & message fields
- **Footer** — CTAs, navigation links, branding, copyright

### `about.html` — About Page
A detailed company profile page. Includes:
- **Hero** — Company tagline and philosophy intro
- **Our Vision** — Animated glowing vision card
- **Mentorship Methods** — Numbered methodology rows (01, 02, 03)
- **Team Grid** — 9-card team member display
- **Roadmap** — "Road to 1 Crore" animated SVG path with markers
- **Segments** — Three glowing pill modules (Regular, Trading, Live Sessions)
- **Footer** — Shared footer component

### `form.html` — Registration Form
Standalone enrollment form page. Includes:
- **Multi-field form** — Name, Place, Contact, Email, Status, Course
- **Trading Background selector** — Radio cards for Beginner / Intermediate / Advanced
- **Privacy Agreement** — Checkbox consent
- **Submit button** with smooth hover effects

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
- 🌑 **Deep dark green-black** background (`#030805`)
- 💚 **Neon green** (`#00ff00`) accent with animated glow effects
- 🔤 **Playfair Display** serif for headings — elegant, premium feel
- 📝 **Inter** sans-serif for body — clean, readable
- ✨ **Smooth scroll-reveal** on all major sections (IntersectionObserver)
- 🌟 **Glassmorphism** & gradient border effects on key cards
- 🎞️ **Infinite marquee** testimonial scroll (dual-direction rows)

---

## ⚙️ Features & Effects

| Feature | Implementation |
|---------|---------------|
| Smooth Scroll | `html { scroll-behavior: smooth; }` |
| Scroll Reveal | `IntersectionObserver` + `.reveal` / `.reveal.active` CSS classes |
| Glow Animations | `@keyframes` with `text-shadow`, `filter: blur`, `box-shadow` |
| Gradient Borders | `::before` pseudo-element with `z-index` stacking |
| Marquee Scroll | CSS `@keyframes marquee-left / marquee-right` on `.marquee-track` |
| Stagger Reveal | `.reveal-stagger` class with `nth-child` transition delays |
| Favicon | `<link rel="icon">` pointing to `assets/logo.png` |

---

## 🚀 Getting Started

No build tools or dependencies required. Just open in a browser:

```bash
# Clone the repository
git clone https://github.com/abhinandpn/Website.git

# Open in browser
open index.html
```

Or simply double-click `index.html` to view locally.

---

## 🔗 Internal Navigation

| From | Link | Destination |
|------|------|-------------|
| `index.html` navbar | Home | `index.html` |
| `index.html` navbar | About | `about.html` |
| `index.html` navbar | Contact Us | `index.html#contact` |
| `about.html` navbar | Home | `index.html` |
| `about.html` footer | Contact | `index.html#contact` |
| All footers | Enroll | `form.html` (intended) |

---

## 📞 Contact Information

- 📱 **Phone:** +91 90377 13791
- 📧 **Email:** stockedhelp@gmail.com

---

## 👨‍💻 Author

**Designed & Developed by** `abhinand_pn`

---

<div align="center">

© STOCK-ED. 2024 · All rights reserved.

</div>
