# webkit-268153
https://bugs.webkit.org/show_bug.cgi?id=268153

## Running

```sh
go run main.go
```

The query params are:

* `p_lang`: primary audio's language
* `p_auto`: primary audio's AUTOSELECT value
* `p_default`: primary audio's DEFAULT value
* `ad_lang`: audio description's language
* `ad_auto`: audio description's AUTOSELECT value
* `ad_default`: audio description's DEFAULT value

```sh
curl 'http://localhost:9999/playlist.m3u8?p_lang=it&p_auto=YES&ad_lang=en&ad_auto=NO'
```

Generated the playlists with:

```sh
ffmpeg -i video.mp4 -c copy -hls_segment_type fmp4 -hls_list_size 0 -hls_time 6 -hls_fmp4_init_filename "v_init.mp4" v_index.m3u8

ffmpeg -i audio.mp4 -c copy -hls_segment_type fmp4 -hls_list_size 0 -hls_time 6 -hls_fmp4_init_filename "a_init.mp4" a_index.m3u8

ffmpeg -i audio_ad.mp4 -c copy -hls_segment_type fmp4 -hls_list_size 0 -hls_time 6 -hls_fmp4_init_filename "ad_init.mp4" ad_index.m3u8
```

## Specs 
https://datatracker.ietf.org/doc/html/draft-pantos-hls-rfc8216bis

> **DEFAULT**  
The value is an enumerated-string; valid strings are YES and NO.
If the value is YES, then the client SHOULD play this Rendition of
the content in the absence of information from the user indicating
a different choice.  
This attribute is OPTIONAL.  Its absence
indicates an implicit value of NO.

> **AUTOSELECT**  
The value is an enumerated-string; valid strings are YES and NO.
This attribute is OPTIONAL.  Its absence indicates an implicit
value of NO.  If the value is YES, then the client MAY choose to
play this Rendition in the absence of explicit user preference
because it matches the current playback environment, such as
chosen system language.  
If the AUTOSELECT attribute is present, its value MUST be YES if
the value of the DEFAULT attribute is YES.

## Test Results

System locale set to "en".
AD has `CHARACTERISTICS="public.accessibility.describes-video"`


Primary: No tags  
AD: No tags  
Actual: Primary  
Expected: Primary  

Primary: No tags  
AD: language="en"  
Actual: AD  
Expected: AD  

Primary: language="en"  
AD: language="en"  
Actual: Primary  
Expected: Primary  

Primary: language="it"  
AD: language="en"  
Actual: AD  
Expected: AD  

Primary: language="it", AUTOSELECT=YES  
AD: language="en"  
Actual: AD  
Expected: Primary  

Primary: language="it", AUTOSELECT=YES  
AD: language="en", AUTOSELECT=NO  
Actual: AD  
Expected: Primary

Primary: language="it", AUTOSELECT=YES, DEFAULT=YES  
AD: language="en", AUTOSELECT=NO  
Actual: AD  
Expected: Primary

Primary: language="en", AUTOSELECT=YES  
AD: language="en", AUTOSELECT=NO  
Actual: Primary  
Expected: Primary 

Primary: language="en", AUTOSELECT=NO  
AD: language="en", AUTOSELECT=YES  
Actual: Primary  
Expected: AD 
