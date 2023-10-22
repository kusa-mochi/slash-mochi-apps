/** @type {import('next').NextConfig} */
const isDev = process.env.NODE_EVN !== 'production'
const nextConfig = {
  output:'export',
  images: {
    unoptimized: true,
  },
  assetPrefix: isDev ? '' : '/',
  reactStrictMode: true,
}

module.exports = nextConfig
