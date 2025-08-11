# Step-by-Step Deployment Guide

## Prerequisites
- GitHub account with your code pushed
- Email for Render/Vercel signup

## Part 1: Deploy Backend to Render (10 minutes)

### Step 1: Sign up for Render
1. Go to https://render.com
2. Click "Get Started for Free"
3. Sign up with GitHub (recommended) or email

### Step 2: Create New Web Service
1. From dashboard, click **"New +"** button
2. Select **"Web Service"**

### Step 3: Connect Your Repository
1. If using GitHub auth: Select **"curriculum-vitae"** from your repos
2. If not: Click **"Public Git repository"** and enter:
   ```
   https://github.com/rockstar89/curriculum-vitae
   ```

### Step 4: Configure Service
Fill in these settings:
- **Name:** `cv-backend` (or any name you prefer)
- **Region:** Choose closest to you (e.g., Oregon USA / Frankfurt EU)
- **Branch:** `master`
- **Root Directory:** `backend`
- **Runtime:** `Docker`
- **Instance Type:** `Free`

### Step 5: Add Environment Variables
Scroll down to **"Environment Variables"** and add:

| Key | Value |
|-----|-------|
| PORT | 8080 |
| GIN_MODE | release |
| ADMIN_USERNAME | *(choose your username)* |
| ADMIN_PASSWORD | *(choose strong password)* |
| JWT_SECRET | *(generate random 32 chars)* |

**To generate JWT_SECRET:** Use https://randomkeygen.com/ (256-bit key)

### Step 6: Deploy
1. Click **"Create Web Service"**
2. Wait 5-10 minutes for build and deploy
3. Once deployed, you'll see: `https://cv-backend.onrender.com`
4. **Save this URL - you'll need it for frontend!**

### Step 7: Test Backend
Visit: `https://YOUR-BACKEND-URL.onrender.com/health`
Should see: `{"service":"cv-backend","status":"ok"}`

---

## Part 2: Deploy Frontend to Vercel (5 minutes)

### Step 1: Sign up for Vercel
1. Go to https://vercel.com
2. Click **"Sign Up"**
3. Continue with GitHub (recommended)

### Step 2: Import Project
1. Click **"Add New..."** → **"Project"**
2. Under "Import Git Repository", find **"curriculum-vitae"**
3. Click **"Import"**

### Step 3: Configure Project
1. **Framework Preset:** Create React App (auto-detected)
2. **Root Directory:** Click "Edit" and change to `frontend`
3. **Build Settings:** (leave defaults)
   - Build Command: `npm run build`
   - Output Directory: `build`

### Step 4: Add Environment Variable
Click **"Environment Variables"** and add:

| Name | Value |
|------|-------|
| REACT_APP_API_URL | *Your Render backend URL* |

Example: `https://cv-backend.onrender.com`

### Step 5: Deploy
1. Click **"Deploy"**
2. Wait 2-3 minutes for build
3. You'll get URL like: `https://curriculum-vitae-xxx.vercel.app`

---

## Part 3: Post-Deployment Setup

### Update CORS (Important!)
1. Go back to Render dashboard
2. Click on your backend service
3. Go to **"Environment"** tab
4. Add new variable:
   ```
   ALLOWED_ORIGINS = https://your-app.vercel.app
   ```
5. Render will auto-redeploy

### Test Everything
1. **Visit your Vercel URL** - CV should load
2. **Test Admin:** Go to `/admin` on your Vercel URL
3. **Login** with credentials you set in Render
4. **Upload a CV** to test file upload
5. **Download CV** from main page

---

## Troubleshooting

### Backend not responding?
- Render free tier sleeps after 15 minutes
- First request takes 5-30 seconds to wake up
- This is normal for free tier

### CORS errors?
- Make sure `ALLOWED_ORIGINS` in Render matches your Vercel URL exactly
- Include `https://` in the URL

### Login not working?
- Check username/password match what you set in Render
- Check browser console for errors
- Ensure backend is running (check /health endpoint)

### File upload fails?
- Render free tier has 512MB disk space
- Check if backend is awake
- Verify JWT_SECRET is set correctly

---

## Custom Domain (Optional)

### For Frontend (Vercel):
1. Go to project settings → Domains
2. Add your domain
3. Follow DNS instructions

### For Backend (Render):
1. Upgrade to paid plan required
2. Or use subdomain like api.yourdomain.com

---

## Monitoring

### Render Dashboard Shows:
- Deploy status
- Logs (last 100 lines free)
- Metrics (basic)
- Service health

### Vercel Dashboard Shows:
- Deployment history
- Function logs
- Analytics (basic)
- Build logs

---

## Total Time: ~20 minutes
## Total Cost: $0/month

Backend will sleep when inactive but wakes up automatically when needed!