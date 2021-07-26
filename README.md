# Golang video generator server

<!-- ABOUT THE PROJECT -->
## About The Project
This is the video generator for [top video generator app](https://github.com/mage1711/video-generator-app) it downloads and generates the videos

### Built With

* [Golang](https://golang.org/)
* [FFmpeg](http://ffmpeg.org/)
* [Cloudinary](https://cloudinary.com/)
* [Docker](https://www.docker.com/)
## How it works

The server waits for the json from the app that contains video name and url and if 
the url is not present the app fetches a url from YouTube.</br>
after the videos are downloaded by the app they use a small package i built over ffmpeg that concatenate videos together and sends the video link.

## Example of sent json
```json
[
    {
        "filename": "Chicory A Colorful Tale",
        "name": "Chicory: A Colorful Tale",
        "platform": "PlayStation 5",
        "playlist_found": false,
        "position": 0,
        "rating": "87",
        "release_date": [
            "June",
            "10",
            "2021"
        ],
        "url": "https://www.metacritic.com/game/playstation-5/chicory-a-colorful-tale",
        "video_found": false
    },
    {
        "filename": "Mass Effect Legendary Edition",
        "name": "Mass Effect Legendary Edition",
        "platform": "PC",
        "playlist_found": true,
        "playlist_videos": [
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/02/01/646528/trailer_masseffectLE_202121_4000.mp4",
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/04/13/649310/trailer_masseffectlegend_comparison_4000.mp4",
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/05/12/649991/MELE1_CharacterCreatorGameplay_8000.mp4",
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/05/13/650014/GraphicsComp_MassEffect1_20210513_8000.mp4"
        ],
        "position": 1,
        "rating": "86",
        "release_date": [
            "May",
            "14",
            "2021"
        ],
        "url": "https://www.metacritic.com/game/pc/mass-effect-legendary-edition",
        "video_found": true,
        "video_url": "https://static-gamespotvideo.cbsistatic.com/vr/2021/02/09/646974/Feature_MELEComp_20210208_8000.mp4"
    },
    {
        "filename": "Final Fantasy VII Remake Intergrade",
        "name": "Final Fantasy VII Remake Intergrade",
        "platform": "PlayStation 5",
        "playlist_found": true,
        "playlist_videos": [
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/03/22/648808/trailer_ff7remakeintegradeextended_4000.mp4",
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/06/09/650625/PS5_FF7RScorpionFight4K_8000.mp4",
            "https://static-gamespotvideo.cbsistatic.com/vr/2021/06/10/650666/Feature_FF7RIGFXComparison_20210609_8000.mp4"
        ],
        "position": 2,
        "rating": "89",
        "release_date": [
            "June",
            "10",
            "2021"
        ],
        "url": "https://www.metacritic.com/game/playstation-5/final-fantasy-vii-remake-intergrade",
        "video_found": true,
        "video_url": "https://static-gamespotvideo.cbsistatic.com/vr/2021/05/17/650075/Trailer_FF7intergrade_finaltrailer_4000.mp4"
    }
]
```


## Example videos
```
https://www.youtube.com/watch?v=BdqRhyQ_QAU&t=6
https://res.cloudinary.com/dvmo50ocz/video/upload/v1627297087/wtatvopyox0u2pqeihjb.mp4
```
