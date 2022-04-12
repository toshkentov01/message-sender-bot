# message-sender-bot
This is message sender bot, which sends messages to channels or groups



HOW TO USE:
    
    1. Create a bot in telegram. 
        You can create a bot with botfather bot.
        Link of this bot is: @BotFather
    
    2. After that botfather will give you the token. Copy this token and set to .env file

    3. Make this bot one of the admin of channel, where you want to send a message. 

    4. Then in .env file also set your channel's username

    5. Then run the program

    -------------------------------------------------------------------------------------

    6. There are two APIs:

        First API sends messages and returns status ok, if there will not be an error

        Second API is sends messages periodically (each 5 second) and will never finish its work

    !!! WARNING !!!
    AFTER CLONNING A message-sender-bot PROJECT, PLEASE OPEN api_gateway IN ONE EDITOR (vscode, goland ...) 
    AND message_service IN ANOTHER EDITOR.

    DON'T FORGET TO SETTING AN .ENV FILE