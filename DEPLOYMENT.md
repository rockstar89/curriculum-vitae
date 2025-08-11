# Deployment Guide

## FREE Deployment Options

### Option 1: Vercel + Render (FREE)

**Frontend (Vercel - FREE)**
- Static React app deployment
- Global CDN and SSL
- Automatic deployments from GitHub
- Custom domain support

**Backend (Render - FREE)**
- 512MB RAM
- Spins down after 15 minutes of inactivity
- Spins up on request (5-30 second delay)
- File storage persists during sleep
- Automatic HTTPS

### Total Cost: $0/month

---

### Option 2: Netlify + Cyclic.sh (FREE)

**Frontend (Netlify - FREE)**
- Static hosting with CDN
- Continuous deployment
- Custom domain support

**Backend (Cyclic.sh - FREE)**
- AWS Lambda-based (serverless)
- 10,000 requests/month free
- Built-in file storage (S3)
- Always-on (no cold starts)
- Automatic HTTPS

### Total Cost: $0/month

---

## Deployment Steps

### 1A. Deploy Backend to Render (FREE)

1. Go to [Render](https://render.com)
2. Sign in with GitHub
3. Click "New +" → "Web Service"
4. Connect your GitHub repository
5. Configure:
   - **Name:** cv-backend
   - **Region:** Choose nearest
   - **Branch:** master
   - **Root Directory:** backend
   - **Runtime:** Docker
   - **Plan:** Free
6. **IMPORTANT:** Add environment variables in Render Dashboard:
   ```
   PORT=8080                              # Already in render.yaml
   GIN_MODE=release                       # Already in render.yaml
   ADMIN_USERNAME=your-username-here     # SET YOUR OWN
   ADMIN_PASSWORD=your-secure-password   # SET YOUR OWN  
   JWT_SECRET=generate-random-256-bit    # SET YOUR OWN
   ```
   **Security Note:** Never commit these values to your repository. Set them directly in Render's environment variables section.
7. Click "Create Web Service"
8. Note your Render URL (e.g., `https://cv-backend.onrender.com`)

### 1B. Deploy Backend to Cyclic.sh (FREE)

1. Go to [Cyclic.sh](https://cyclic.sh)
2. Sign in with GitHub
3. Click "Deploy" → Select your repository
4. Cyclic will auto-detect the Go app
5. Set environment variables in dashboard:
   ```
   PORT=8080
   GIN_MODE=release
   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=your-secure-password
   JWT_SECRET=your-jwt-secret-key
   ```
6. Deploy and get your URL (e.g., `https://your-app.cyclic.app`)

### 2. Deploy Frontend to Vercel

1. Go to [Vercel](https://vercel.com)
2. Sign in with GitHub
3. Click "New Project"
4. Import your `curriculum-vitae` repository
5. Vercel will auto-detect React app in `/frontend`
6. Set environment variable:
   ```
   REACT_APP_API_URL=https://your-railway-app.railway.app
   ```
7. Deploy

### 3. Update CORS Settings

After getting your Vercel URL, update backend CORS to allow your domain:

In `backend/internal/middleware/cors.go`, add your Vercel domain to allowed origins.

---

## Paid Alternatives (Better Performance)

### Option 3: Railway (~$5/month)
- **Frontend:** Vercel (FREE)
- **Backend:** Railway Hobby ($5/month)
- Always-on, no cold starts
- Better for high-traffic sites

### Option 4: Single VPS (~$5/month)
- DigitalOcean/Linode/Vultr droplet
- Run both services with Docker Compose
- Manual setup but full control

## Free Platform Comparison

| Platform | Cold Start | Storage | Limits | Best For |
|----------|------------|---------|---------|----------|
| Render | 5-30s after 15min | Persistent | 750hrs/month | Low-traffic portfolios |
| Cyclic.sh | None (serverless) | S3 included | 10k requests/month | Consistent traffic |
| Fly.io | None (3 machines) | Persistent volumes | 3 VMs, 160GB transfer | Production apps |
| Glitch | 5s after 5min | Persistent | Always-on with boosting | Development/demos |

---

## Environment Variables Reference

### Backend (Render/Cyclic/Railway)
```bash
PORT=8080
GIN_MODE=release
ADMIN_USERNAME=admin
ADMIN_PASSWORD=secure-password-here
JWT_SECRET=your-jwt-secret-256-bit
ALLOWED_ORIGINS=https://your-vercel-app.vercel.app
```

### Frontend (Vercel/Netlify)
```bash
# For Render backend:
REACT_APP_API_URL=https://your-app.onrender.com
# For Cyclic backend:
REACT_APP_API_URL=https://your-app.cyclic.app
REACT_APP_ENV=production
```

---

## Post-Deployment Checklist

- [ ] Frontend loads correctly
- [ ] Language switching works
- [ ] CV download works (public)
- [ ] Admin login works
- [ ] CV upload works (admin)
- [ ] Password change works (admin)
- [ ] File persistence works
- [ ] HTTPS enabled on both services

---

## Monitoring & Maintenance

**Render (Free):**
- Monitor spin-up times in logs
- Check monthly usage (750 hours limit)
- Upgrade to paid if needed ($7/month starter)

**Cyclic (Free):**
- Monitor request count (10k/month limit)
- Check S3 storage usage
- Upgrade if exceeding limits

**Vercel:**
- Monitor build logs
- Automatic deployments from GitHub
- Analytics available in dashboard

---

## Cost Breakdown

| Service | Plan | Cost/Month | Features | Limitations |
|---------|------|------------|----------|-------------|
| Vercel | Hobby | FREE | Static hosting, CDN, SSL | None for static sites |
| Render | Free | FREE | 512MB RAM, Auto-SSL | Spins down after 15min |
| Cyclic | Free | FREE | Serverless, S3 storage | 10k requests/month |
| Netlify | Free | FREE | Static hosting, CDN | 100GB bandwidth/month |
| **Total** | | **$0/month** | Full-stack deployment | Some cold starts |

**Note:** Free tiers work perfectly for portfolios and low-traffic sites. The main tradeoff is potential cold starts (5-30 seconds) when the backend wakes up after inactivity.