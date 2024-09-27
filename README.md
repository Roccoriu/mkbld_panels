# MKBLD (because grifting is bad)

This is a cross-platform image downloader written in Go. It can download the images from the panels app for free and in full quality. It downloads everything it can find and sorts it appropriately
## Features

- Supports downloading multiple image qualities.
- Automatic sorting based on image quality as specified by the Web API


## Downloads

Here you can find the relevant binaries for each platform. ARM binaries are included for all three platforms. 

- [Windows (amd64)](https://github.com/Roccoriu/mkbld_panels/releases/download/v0.1.0/mkbhd_downloader_windows_amd64.exe)
- [Windows (arm64)](https://github.com/Roccoriu/mkbld_panels/releases/download/v0.1.0/mkbhd_downloader_windows_arm64.exe)
- [Linux (amd64)](https://github.com/Roccoriu/mkbld_panels/releases/download/v0.1.0/mkbhd_downloader_linux_amd64)
- [Linux (arm64)](https://github.com/Roccoriu/mkbld_panels/releases/download/v0.1.0/mkbhd_downloader_linux_arm64)
- [MacOS (arm64)](https://github.com/Roccoriu/mkbld_panels/releases/download/v0.1.0/mkbhd_downloader_macos_arm64)


## Run
1. Download the appropriate version for your platform and architecture.
2. Run the application like any other, you may have to allow list it as it is not an officially signed app.
3. Wait until the app finishes, it will automatically close after it is done downloading everything. It will log every step of the process.


## FAQ

### Q: What's the story behind this?

On September 24th, 2024, well-known tech YouTuber MKBHD released Panels, a wallpaper app that:

- Had insanely invasive, unjustified tracking including for location history and search history.
- Charged artists a predatory 50% commission (even Apple takes only 30% for app purchases).
- Forced you to watch two ads for every wallpaper that you wanted to download, and then only letting you download it in SD.
- Gatekept all HD wallpapers behind a **fifty dollars a year subscription**.
- Had many wallpapers that were essentially AI-generated slop or badly edited stock photos.

Especially given MKBHD's previous criticism of substandard companies and products, people justifiably got upset given that this looked like a pretty blatant grift and cash-grab that is exploitative of the fan base that's trusted his editorial integrity over the past fifteen years. However, on the same day, MKBHD wrote a post doubling down on the app.

### Q: Aren't you stealing from artists by running this script?

MKBSD accesses publicly available media through the Panels app's own API. It doesn't do anything shady or illegal. The real problem here is Panels and MKBHD's complete inability to provide a secure platform for the artists that they're ~~exploiting~~ working with. Any other app could have avoided the issues that make MKBSD possible had it been engineered competently.


## License

This project is licensed under the GPLv3 License. See the LICENSE file for more info.
