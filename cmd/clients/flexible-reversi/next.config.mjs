/** @type {import('next').NextConfig} */

const isDevelopment = process.env.NODE_ENV !== 'production'
const nextConfig = {
    output:'export',
    images: {
        unoptimized: true
    },
    assetPrefix: isDevelopment ? '' : '/flexible-reversi/',
};

export default nextConfig;
