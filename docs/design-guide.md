# DevHub Design System Documentation

A comprehensive guide to recreating the sleek, modern dark theme UI aesthetic used in DevHub.

---

## 🎨 Design Philosophy

**Visual Identity**: Cyberpunk-inspired dark theme with glassmorphism, neon accents, and professional polish
**Target Aesthetic**: Modern developer tools (VS Code, Vercel, Linear, Raycast)
**Key Principles**:
- High contrast for readability
- Subtle depth through layering and shadows
- Smooth animations and transitions
- Clean, minimalist layout with breathing room

---

## 🎭 Color Palette

### Background Colors
```css
/* Primary Backgrounds */
--background-primary: #1a1d23      /* Main app background - Deep charcoal */
--background-secondary: #2d3139    /* Card backgrounds - Medium gray */
--background-tertiary: #363a45     /* Hover states, elevated elements */

/* Extended Grays */
--gray-950: #0a0c10               /* Darkest - overlays, shadows */
--gray-900: #13151a               /* Very dark - deep backgrounds */
--gray-800: #1a1d23               /* Main background */
--gray-700: #2d3139               /* Secondary surfaces */
--gray-600: #363a45               /* Tertiary surfaces */
--gray-500: #6b7280               /* Medium gray - inactive text */
--gray-400: #9ca3af               /* Light gray - secondary text */
```

### Accent Colors
```css
/* Primary Accent - Cyan (Electric Blue) */
--cyan: #00D9FF                   /* Primary actions, links, focus states */
--cyan-dark: #00B8D4              /* Hover states */
--cyan-light: #4DFFFF             /* Highlights, glow effects */

/* Secondary Accent - Mint (Seafoam Green) */
--mint: #7FE0C3                   /* Success states, secondary CTAs */
--mint-dark: #5BC4A6              /* Hover states */
--mint-light: #9FEDCE             /* Highlights */
```

### Semantic Colors
```css
/* Status Colors */
--success: #22c55e                /* Green - success states */
--warning: #eab308                /* Yellow - warning states */
--error: #ef4444                  /* Red - error, danger states */
--info: #00D9FF                   /* Cyan - informational */

/* Status with transparency (for badges, pills) */
--success-bg: rgba(34, 197, 94, 0.2)
--success-border: rgba(34, 197, 94, 0.3)
--success-text: #22c55e

--error-bg: rgba(239, 68, 68, 0.2)
--error-border: rgba(239, 68, 68, 0.3)
--error-text: #ef4444
```

### Text Colors
```css
--text-primary: #f9fafb           /* Headings, important text - Near white */
--text-secondary: #9ca3af         /* Body text, descriptions - Light gray */
--text-tertiary: #6b7280          /* Metadata, timestamps - Medium gray */
--text-disabled: #4b5563          /* Disabled states - Dark gray */
```

---

## ✍️ Typography

### Font Family
```css
/* Primary Font: Noto Sans JP */
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@300;400;500;600;700&display=swap');

font-family: "Noto Sans JP", system-ui, sans-serif;
```

**Why Noto Sans JP?**
- Clean, modern sans-serif with excellent readability
- Professional appearance for developer tools
- Wide character support (Latin + Japanese)
- Geometric letterforms with subtle humanist touch

**Alternative fonts** (if you want different options):
- `Inter` - Very popular for modern UIs
- `IBM Plex Sans` - Developer-focused, slightly wider
- `Outfit` - Geometric, modern feel
- `Space Grotesk` - Tech-forward aesthetic

### Font Sizes & Weights

```css
/* Headings */
h1: 36px, font-weight: 700 (bold)
h2: 30px, font-weight: 600 (semibold)
h3: 24px, font-weight: 600 (semibold)
h4: 20px, font-weight: 600 (semibold)
h5: 18px, font-weight: 500 (medium)

/* Body Text */
base: 16px, font-weight: 400 (regular)
small: 14px, font-weight: 400 (regular)
xs: 12px, font-weight: 400 (regular)

/* UI Elements */
button: 14px, font-weight: 500 (medium)
badge: 12px, font-weight: 500 (medium)
label: 14px, font-weight: 500 (medium)
```

### Text Rendering
```css
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
```

---

## 🧊 Glassmorphism Effects

The signature "glass" aesthetic throughout DevHub:

### Base Glass Effect
```css
.glass {
  background: rgba(45, 49, 57, 0.4);        /* Semi-transparent background */
  backdrop-filter: blur(12px);               /* Blur background content */
  -webkit-backdrop-filter: blur(12px);       /* Safari support */
  border: 1px solid rgba(107, 114, 128, 0.2); /* Subtle border */
}
```

### Glass Card Component
```css
.glass-card {
  background: rgba(45, 49, 57, 0.4);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(107, 114, 128, 0.2);
  border-radius: 12px;                       /* Rounded corners */
  padding: 24px;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.4),
              0 4px 6px -4px rgba(0, 0, 0, 0.4);
}
```

### Glass Variations
```css
/* Intense Glass (more opacity) */
.glass-intense {
  background: rgba(45, 49, 57, 0.7);
  backdrop-filter: blur(16px);
}

/* Light Glass (less opacity) */
.glass-light {
  background: rgba(45, 49, 57, 0.2);
  backdrop-filter: blur(8px);
}
```

---

## 📦 Component Patterns

### Buttons

#### Primary Button (Cyan Accent)
```css
.btn-primary {
  background: #00D9FF;
  color: #1a1d23;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 500;
  font-size: 14px;
  transition: all 0.2s ease;
  border: none;
}

.btn-primary:hover {
  background: #4DFFFF;
  transform: translateY(-1px);
  box-shadow: 0 8px 16px rgba(0, 217, 255, 0.3);
}
```

#### Secondary Button (Glass Style)
```css
.btn-secondary {
  background: rgba(45, 49, 57, 0.4);
  backdrop-filter: blur(12px);
  color: #f9fafb;
  border: 1px solid rgba(107, 114, 128, 0.2);
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 500;
  font-size: 14px;
  transition: all 0.2s ease;
}

.btn-secondary:hover {
  background: rgba(54, 58, 69, 0.6);
  border-color: rgba(0, 217, 255, 0.4);
}
```

#### Icon Button
```css
.btn-icon {
  background: transparent;
  color: #9ca3af;
  padding: 8px;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.btn-icon:hover {
  color: #00D9FF;
  background: rgba(0, 217, 255, 0.1);
}
```

### Cards

#### Standard Card
```css
.card {
  background: rgba(45, 49, 57, 0.4);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(107, 114, 128, 0.2);
  border-radius: 12px;
  padding: 24px;
  transition: all 0.3s ease;
}

.card:hover {
  border-color: rgba(0, 217, 255, 0.4);
  transform: translateY(-2px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.5);
}
```

#### Info Card (Smaller stats cards)
```css
.info-card {
  background: rgba(45, 49, 57, 0.4);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(107, 114, 128, 0.2);
  border-radius: 10px;
  padding: 16px;
}
```

### Badges & Pills

#### Status Badge
```css
.badge {
  padding: 4px 12px;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 500;
  border: 1px solid;
}

/* Success variant */
.badge-success {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
  border-color: rgba(34, 197, 94, 0.3);
}

/* Error variant */
.badge-error {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
  border-color: rgba(239, 68, 68, 0.3);
}

/* Cyan variant */
.badge-cyan {
  background: rgba(0, 217, 255, 0.2);
  color: #00D9FF;
  border-color: rgba(0, 217, 255, 0.3);
}
```

### Input Fields

```css
.input {
  background: rgba(26, 29, 35, 0.6);
  border: 1px solid rgba(107, 114, 128, 0.2);
  border-radius: 8px;
  padding: 10px 14px;
  color: #f9fafb;
  font-size: 14px;
  transition: all 0.2s ease;
}

.input:focus {
  outline: none;
  border-color: #00D9FF;
  box-shadow: 0 0 0 3px rgba(0, 217, 255, 0.15);
}

.input::placeholder {
  color: #6b7280;
}
```

### Tabs

```css
.tabs-container {
  background: rgba(45, 49, 57, 0.4);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(107, 114, 128, 0.2);
  border-radius: 10px;
  padding: 4px;
  display: flex;
  gap: 4px;
}

.tab {
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  color: #9ca3af;
  transition: all 0.2s ease;
  background: transparent;
}

.tab:hover {
  color: #f9fafb;
}

.tab.active {
  background: #00D9FF;
  color: #1a1d23;
}
```

---

## 🎬 Animations & Transitions

### Standard Transitions
```css
/* All interactive elements */
transition: all 0.2s ease;

/* Hover lift effect */
transition: transform 0.2s ease, box-shadow 0.2s ease;
transform: translateY(-2px);

/* Fade in */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

/* Pulse (loading states) */
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

/* Slide up */
@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
```

### Timing Functions
```css
ease: cubic-bezier(0.25, 0.1, 0.25, 1.0)     /* Standard easing */
ease-out: cubic-bezier(0, 0, 0.2, 1)          /* Decelerate */
ease-in-out: cubic-bezier(0.4, 0, 0.2, 1)    /* Smooth both ends */
```

---

## 📏 Spacing & Layout

### Spacing Scale
```css
xs: 4px
sm: 8px
md: 16px
lg: 24px
xl: 32px
2xl: 48px
3xl: 64px
```

### Border Radius
```css
sm: 6px     /* Small elements, badges */
md: 8px     /* Buttons, inputs */
lg: 10px    /* Info cards */
xl: 12px    /* Cards, dialogs */
2xl: 16px   /* Large containers */
full: 9999px /* Circular */
```

### Shadows
```css
/* Card shadow */
box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.4),
            0 4px 6px -4px rgba(0, 0, 0, 0.4);

/* Hover shadow (lifted) */
box-shadow: 0 12px 24px rgba(0, 0, 0, 0.5),
            0 8px 16px rgba(0, 0, 0, 0.3);

/* Glow effect (cyan) */
box-shadow: 0 0 20px rgba(0, 217, 255, 0.3);
```

---

## 🎯 Icons

**Icon Library**: Lucide React
**Size**: 16px (w-4 h-4), 20px (w-5 h-5), 24px (w-6 h-6)
**Style**: Outlined, stroke width 2

### Common Icons Used
```typescript
import {
  // Navigation
  ArrowLeft, ArrowRight, ChevronLeft, ChevronRight,

  // Actions
  Plus, Minus, X, Check, Copy, Edit, Trash2,
  MoreVertical, MoreHorizontal,

  // File & Folder
  Folder, File, FileText, FolderOpen,

  // Dev Tools
  Terminal, Code2, GitBranch, Package, Container,

  // UI
  Search, Filter, LayoutGrid, List,

  // Status
  AlertCircle, CheckCircle, Clock, Activity,

  // Data
  Database, HardDrive, Cpu, Zap
} from 'lucide-react';
```

---

## 📜 Scrollbar Styling

Custom auto-hide scrollbar that appears on hover:

```css
/* Base scrollbar (hidden by default) */
* {
  scrollbar-width: thin;
  scrollbar-color: transparent transparent;
}

/* Show on hover */
*:hover {
  scrollbar-color: rgba(100, 116, 139, 0.5) transparent;
}

/* Webkit (Chrome, Safari, Edge) */
*::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

*::-webkit-scrollbar-track {
  background: transparent;
}

*::-webkit-scrollbar-thumb {
  background: transparent;
  border-radius: 4px;
}

*:hover::-webkit-scrollbar-thumb {
  background: rgba(100, 116, 139, 0.5);
}

*::-webkit-scrollbar-thumb:hover {
  background: rgba(100, 116, 139, 0.7);
}
```

---

## 🛠️ Implementation Guide

### For Tailwind CSS Projects

**1. Install Tailwind and dependencies:**
```bash
npm install -D tailwindcss postcss autoprefixer @tailwindcss/typography
npx tailwindcss init -p
```

**2. Configure `tailwind.config.js`:**
```javascript
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        background: {
          DEFAULT: '#1a1d23',
          secondary: '#2d3139',
          tertiary: '#363a45',
        },
        cyan: {
          DEFAULT: '#00D9FF',
          dark: '#00B8D4',
          light: '#4DFFFF',
        },
        mint: {
          DEFAULT: '#7FE0C3',
          dark: '#5BC4A6',
          light: '#9FEDCE',
        },
        gray: {
          950: '#0a0c10',
          900: '#13151a',
          800: '#1a1d23',
          700: '#2d3139',
          600: '#363a45',
          500: '#6b7280',
          400: '#9ca3af',
        },
      },
      fontFamily: {
        sans: ['"Noto Sans JP"', 'system-ui', 'sans-serif'],
      },
      backdropBlur: {
        xs: '2px',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
}
```

**3. Add to your CSS (`index.css` or `globals.css`):**
```css
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@300;400;500;600;700&display=swap');

@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  body {
    @apply bg-background text-gray-100 font-sans antialiased;
  }

  /* Custom scrollbar styles (see above) */
}

@layer utilities {
  .glass {
    @apply bg-background-secondary/40 backdrop-blur-md border border-gray-600/20;
  }

  .glass-card {
    @apply glass rounded-xl p-6 shadow-lg;
  }
}
```

### For Plain CSS Projects

Create a `variables.css`:
```css
:root {
  /* Colors */
  --bg-primary: #1a1d23;
  --bg-secondary: #2d3139;
  --bg-tertiary: #363a45;

  --cyan: #00D9FF;
  --cyan-dark: #00B8D4;
  --cyan-light: #4DFFFF;

  --text-primary: #f9fafb;
  --text-secondary: #9ca3af;
  --text-tertiary: #6b7280;

  /* Spacing */
  --spacing-xs: 4px;
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --spacing-lg: 24px;
  --spacing-xl: 32px;

  /* Border Radius */
  --radius-sm: 6px;
  --radius-md: 8px;
  --radius-lg: 12px;
}

body {
  background: var(--bg-primary);
  color: var(--text-primary);
  font-family: "Noto Sans JP", system-ui, sans-serif;
  -webkit-font-smoothing: antialiased;
}
```

---

## 💬 Prompt Template for AI Assistants

Use this prompt when starting a new project:

```
Create a modern web application with a sleek dark theme design inspired by developer tools like VS Code and Vercel. Use the following design system:

**Color Palette:**
- Background: #1a1d23 (primary), #2d3139 (secondary), #363a45 (tertiary)
- Primary accent: Cyan #00D9FF
- Secondary accent: Mint #7FE0C3
- Text: Near-white #f9fafb (headings), #9ca3af (body)

**Typography:**
- Font: Noto Sans JP (weights: 300, 400, 500, 600, 700)
- Headings: Bold (700), sizes from 36px to 18px
- Body: Regular (400), 16px base size
- Apply antialiasing

**UI Style:**
- Glassmorphism effects: Semi-transparent backgrounds (rgba(45, 49, 57, 0.4)) with backdrop-blur (12px)
- Rounded corners: 8-12px for cards, 6px for buttons
- Subtle borders: 1px solid rgba(107, 114, 128, 0.2)
- Soft shadows: Use multiple layers for depth
- Smooth transitions: 0.2s ease for all interactive elements

**Components:**
- Cards: Glass effect with hover lift (translateY(-2px))
- Buttons: Primary (cyan bg, dark text), Secondary (glass with border)
- Badges: Status colors with 20% opacity backgrounds
- Icons: Lucide React, outlined style, 20px default size

**Interactions:**
- Hover states: Brighten colors, add glow effects, lift elements
- Focus states: Cyan border with soft shadow ring
- Loading states: Pulse animation on cyan text
- Auto-hide scrollbars that appear on hover

**Layout:**
- 8px spacing scale (4, 8, 16, 24, 32, 48, 64)
- Max content width: 1280px (max-w-7xl)
- Generous padding: 32px (p-8) on main containers

Please implement this design system consistently across all components.
```

---

## 📸 Visual Examples

### Component Hierarchy
```
Dark Background (#1a1d23)
└─ Glass Card (background-secondary/40 + blur)
   ├─ Heading (text-white, bold)
   ├─ Description (text-gray-400)
   ├─ Cyan Button (primary action)
   └─ Badge (status with colored background)
```

### Color Usage Guidelines
- **Cyan**: Primary actions, links, active states, focus indicators
- **Mint**: Success messages, positive confirmations, secondary CTAs
- **Gray scale**: Backgrounds (darker = deeper), text (lighter = more prominent)
- **Red (#ef4444)**: Errors, destructive actions, alerts
- **Green (#22c55e)**: Success states, completed actions
- **Yellow (#eab308)**: Warnings, pending states

---

## 🎓 Best Practices

1. **Maintain contrast ratios**: Aim for WCAG AA (4.5:1 for normal text)
2. **Use glassmorphism sparingly**: Elevate important UI, don't overuse
3. **Consistent spacing**: Stick to 8px grid system
4. **Smooth interactions**: All hover/focus states should have transitions
5. **Icon sizes**: Keep uniform (20px or 16px) within component sections
6. **Loading states**: Use cyan color with pulse animation
7. **Error states**: Red with subtle background, clear messaging
8. **Responsive design**: Cards adapt to grid (4 cols → 2 cols → 1 col)

---

## 📦 Quick Start Checklist

- [ ] Install Noto Sans JP font
- [ ] Set up Tailwind with custom colors
- [ ] Add glassmorphism utility classes
- [ ] Implement custom scrollbar styles
- [ ] Install Lucide React for icons
- [ ] Set up transition defaults
- [ ] Create reusable button components
- [ ] Create card component variants
- [ ] Test color contrast ratios
- [ ] Verify responsive breakpoints

---

**Created for DevHub Design System**
Version 1.0 | October 2025

This design system emphasizes modern developer tool aesthetics with a focus on dark themes, glassmorphism, and neon cyan accents. It's optimized for professional applications requiring both visual appeal and functional clarity.
