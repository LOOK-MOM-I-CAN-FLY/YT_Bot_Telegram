# YT_Bot_Telegram


# **Description**

This project is a Telegram bot written in Go that allows users to download and share YouTube videos. The bot provides a simple interface for interacting with users and processing their requests. It uses the Telegram Bot API and integrates YouTube video downloading capabilities.

--------------------------------------------------------------------------------------------------------------------------------------------------------

# **Features**

YouTube Video Downloading: Users can send a YouTube link, and the bot downloads the video.

Telegram Integration: Automatically sends the downloaded video back to the user via Telegram.

Efficient Design: Handles multiple requests efficiently using Go's concurrency features.

Customizable Settings: Configure video quality and file size limits.

Error Handling: Provides clear feedback to users in case of invalid links or errors.

--------------------------------------------------------------------------------------------------------------------------------------------------------

# **Requirements**

To build and run this project, you need:

Go: Version 1.18 or newer.

Telegram Bot Token: Obtain one from BotFather.

YouTube Downloader: A library or tool for downloading YouTube videos, such as yt-dlp or a Go-based alternative.

----------------------------------------------------------------------------

# **Installation**

1. Clone the Repository

```
git clone https://github.com/LOOK-MOM-I-CAN-FLY/YT_Bot_Telegram.git
cd YT_Bot_Telegram
```

2. Set Up Environment Variables

3. Create a .env file in the project root and add your Telegram bot token:

```TELEGRAM_BOT_TOKEN=your_bot_token_here```

4. Install Dependencies

Use go mod to install required libraries:

```go mod tidy```

5. Run the Bot

Execute the bot using the Go command:

```go run main.go```

The bot will start and listen for incoming messages.

----------------------------------------------------------------------------

# **Usage**

Start the Bot: Send /start to initialize the conversation.

Download Videos:

Send a YouTube link to the bot.

The bot will process the link, download the video, and send it back to you.

Error Handling: If a link is invalid, the bot will notify you and provide guidance.

----------------------------------------------------------------------------

# **File Structure**

main.go: Entry point of the application.

bot/: Contains the core bot logic and Telegram API integration.

downloader/: Handles video downloading and processing.

utils/: Utility functions and helpers.

----------------------------------------------------------------------------

# **Contributing**

Contributions are welcome! If you want to contribute:

1. Fork the repository.

2. Create a new branch for your feature or bug fix:

```git checkout -b feature-name```

3. Commit your changes:

```git commit -m "Description of changes"```

4. Push your branch:

```git push origin feature-name```

5. Create a Pull Request.

----------------------------------------------------------------------------

# **Future Plans**

Support for audio-only downloads.

Add user authentication and rate limiting.

Implement a database for tracking user activity.

Provide options for different video qualities and formats.

Enhance error reporting and logging.

----------------------------------------------------------------------------

# **Acknowledgments**

Telegram Bot API: For enabling seamless bot interactions.

Open Source Community: For tools and libraries that make this project possible.

YouTube: For providing an endless source of video content.

----------------------------------------------------------------------------

# **Contact**

If you have any questions or suggestions, feel free to reach out:

GitHub Issues: Submit your issues here.

Email: igrik315.nekrasov@yandex.ru

Telegram: @Sindi_hall

----------------------------------------------------------------------------

Enjoy using this bot and feel free to contribute to its development!

