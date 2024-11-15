## Bot Package

### Overview
This Go-based application functions as a Discord bot designed to:
* **Scrape data:** Periodically collects information from a specified source.
* **Send to Discord:** Post the scraped data to a designated channel on a Discord server.

### Code Breakdown
#### Package `bot`
* **BotHandler interface:** Defines a contract for any object that can send a message.
* **Bot struct:** Represents the core bot instance, containing:
  * `token`: Discord bot token for authentication.
  * `timeReload`: Interval in minutes for data refresh and message sending.
  * `channel_id`: ID of the Discord channel where messages are sent.
* **SendMessage method:**
  * Connects to Discord using the provided token.
  * Reads data from a file.
  * Sends the data to the specified channel.
  * Initiates a goroutine to repeat the process every `timeReload` minutes.

#### Key Functions and Components
* **`data.NewData()`:** Creates a new instance of the data handler, responsible for managing data files.
* **`scrap.NewScrap()`:** Creates a new instance of the scraper, responsible for fetching data.
* **`discordgo.New()`:** Creates a new Discord session.

### Configuration
* **Discord Bot Token:** Replace `"Bot MTMwNDU5Mzg3ODQ1MTYyMTkxOA.G3Zl5P.oH3BQ-ZpoqMLj_y5RthniuRh0NwY7ulkndIVWo"` with your actual Discord bot token.
* **Channel ID:** Replace `"1304593207732080685"` with the ID of the target Discord channel.
* **Data Source:** The specific data source being scraped is not explicitly defined in the provided code.
* **Data Format:** The format of the data being scraped and stored is also not specified.

### Usage
1. **Build:** Compile the Go code into an executable.
2. **Run:** Execute the binary.
3. **Configuration:** Ensure the Discord bot token and channel ID are set correctly.

### Usage Example 

    //main.go
    package main

    import (
      "scrap.com/bot" // Replace with your project path
      "time"
    )
    func main() {
      // Replace with your actual Discord bot token
      token := "YOUR_BOT_TOKEN" 

      // Replace with the ID of your target Discord channel
      channelID := "YOUR_CHANNEL_ID"
      
      // Create a new bot instance
      myBot := bot.NewBot(token, channelID, 10) // Adjust refresh interval (minutes)
      
      // Start the bot
      go myBot.SendMessage()
      
      // Keep the main program running
      for range time.Tick(time.Second) {}
    }