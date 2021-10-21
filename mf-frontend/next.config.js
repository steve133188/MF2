module.exports = {
  reactStrictMode: true,
}

const path = require('path')
const withSass = require('@zeit/next-sass');
module.exports = withSass({
  /* bydefault config  option Read For More Optios
  here https://github.com/vercel/next-plugins/tree/master/packages/next-sass*/
  cssModules: true
})
module.exports = {
  /* Add Your Scss File Folder Path Here */
  sassOptions: {
    includePaths: [path.join(__dirname, 'styles')],
  },
}
module.exports = {
  webpack5: true,
  webpack: (config) => {
    config.resolve.fallback = { fs: false ,crypto:false , stream: false , path :false};

    return config;
  },
};