# Step-by-Step Deployment Guide

## Prerequisites
- GitHub account with your code pushed
- Email for Render/Vercel signup

## Part 1: Deploy Backend to Render (15 minutes)

### Step 1: Sign up for Render
1. Go to https://render.com
2. Click "Get Started for Free" 
3. Sign up with GitHub (recommended) or email

### Step 2: Create PostgreSQL Database FIRST
⚠️ **Important: Create database before web service!**

1. From dashboard, click **"New +"** button
2. Select **"PostgreSQL"**
3. Configure database:
   - **Name:** `cv-postgres`
   - **Database:** `curriculum_vitae`
   - **User:** `cvadmin`
   - **Region:** Choose closest to you (e.g., Oregon USA / Frankfurt EU)
   - **Plan:** `Free`
4. Click **"Create Database"**
5. **Wait 2-3 minutes** for database to be created
6. **Copy the "Internal Database URL"** - you'll need this next!

### Step 3: Create Web Service
1. Click **"New +"** button → **"Web Service"**
2. Connect your repository:
   - If using GitHub auth: Select **"curriculum-vitae"** from your repos
   - If not: Click **"Public Git repository"** and enter:
     ```
     https://github.com/rockstar89/curriculum-vitae
     ```

### Step 4: Configure Web Service
Fill in these settings:
- **Name:** `cv-backend` (or any name you prefer)
- **Region:** **Same region as your database!**
- **Branch:** `master`
- **Root Directory:** `backend`
- **Runtime:** `Docker`
- **Instance Type:** `Free`

### Step 5: Add Environment Variables
⚠️ **Critical: Add DATABASE_URL first!**

Scroll down to **"Environment Variables"** and add:

| Key | Value | Example |
|-----|-------|---------|
| **DATABASE_URL** | *(Paste Internal Database URL from Step 2)* | `postgresql://cvadmin:pass@dpg-xxx-a.oregon-postgres.render.com/curriculum_vitae` |
| PORT | 8080 | `8080` |
| GIN_MODE | release | `release` |
| ADMIN_USERNAME | *(choose your username)* | `admin` |
| ADMIN_PASSWORD | *(choose strong password)* | `SecurePass123!` |
| JWT_SECRET | *(generate random 32 chars)* | `abc123xyz...` |

**To generate JWT_SECRET:** Use https://randomkeygen.com/ (256-bit key)

### Step 6: Deploy
1. Click **"Create Web Service"**
2. Wait 5-10 minutes for build and deploy
3. Watch the logs - you should see:
   ```
   Attempting to connect to database...
   ✅ Database connected successfully
   🚀 Server starting on port 8080
   ```
4. Once deployed, you'll see: `https://cv-backend.onrender.com`
5. **Save this URL - you'll need it for frontend!**

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
2. Click on your backend service (`cv-backend`)
3. Go to **"Environment"** tab
4. Add new variable:
   - **Key:** `CORS_ORIGIN`
   - **Value:** `https://your-app.vercel.app`
5. Click **"Save Changes"** - Render will auto-redeploy

### Test Everything
1. **Visit your Vercel URL** - CV should load
2. **Test Admin:** Go to `/admin` on your Vercel URL
3. **Login** with credentials you set in Render
4. **Upload a CV** to test file upload (stored in PostgreSQL)
5. **Download CV** from main page

---

## Troubleshooting

### Database Connection Errors?
**Error:** `DATABASE_URL environment variable is required`
**Fix:**
1. Ensure PostgreSQL database was created first
2. Copy the **Internal Database URL** (not External)
3. Add as `DATABASE_URL` environment variable in web service
4. Redeploy

### Backend Build Fails?
**Error:** `go build exit code: 1`
**Fix:** Code has been updated to fix this - pull latest changes

### Backend not responding?
- Render free tier sleeps after 15 minutes
- First request takes 5-30 seconds to wake up
- This is normal for free tier
- Database persists through sleep cycles! ✅

### CORS errors?
- Make sure `CORS_ORIGIN` in Render matches your Vercel URL exactly
- Include `https://` in the URL

### Login not working?
- Check username/password match what you set in Render
- Check browser console for errors
- Ensure backend is running (check /health endpoint)

### File upload fails?
- Files are stored in PostgreSQL database (not filesystem)
- Check if backend is awake
- Verify JWT_SECRET is set correctly
- Database storage has no hibernation issues! ✅

---

## Database Storage Benefits ✅

### Why PostgreSQL Storage?
- **Persistent:** Files survive server hibernation
- **Free:** Uses Render's free 512MB PostgreSQL
- **Reliable:** Database transactions ensure data integrity
- **Scalable:** No filesystem limitations
- **Backup-friendly:** Single database to backup

### File Storage Details
- CV files stored as BYTEA (binary data) in PostgreSQL
- User authentication data in database
- No filesystem dependencies
- Perfect for cloud deployments

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
- Deploy status and logs
- Database connection status
- Service health checks
- PostgreSQL database metrics

### Vercel Dashboard Shows:
- Deployment history
- Function logs
- Analytics (basic)
- Build logs

---

## Architecture Summary

```
Frontend (Vercel) → Backend (Render) → PostgreSQL (Render)
    React/TS          Go/Gin         Database Storage
    
- CV files stored as BYTEA in PostgreSQL
- User data in database tables
- No filesystem dependencies
- $0 cost with persistent storage
```

---

## Total Time: ~20 minutes
## Total Cost: $0/month

**Key Success:** Database storage means no file loss during hibernation! 🎉