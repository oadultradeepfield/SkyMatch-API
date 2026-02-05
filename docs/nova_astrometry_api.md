# Nova.astrometry.net API

The API uses JSON to encode parameters and results. Code documentation is available
in [client.py](https://github.com/dstndstn/astrometry.net/blob/master/net/client/client.py).

## JSON Encoding

First, log in using your API key:

```json
{"apikey": "XXXXXX"}
```

Then encode as `x-www-form-urlencoded`:

```
request-json=%7B%22apikey%22%3A+%22XXXXXX%22%7D
```

Send to `http://nova.astrometry.net/api/login` as a POST request.

### Example with requests library

```python
import requests
import json

R = requests.post('http://nova.astrometry.net/api/login', data={'request-json': json.dumps({"apikey": "XXXXXXXX"})})
print(R.text)
# Output: '{"status": "success", "message": "authenticated user: ", "session": "0ps9ztf2kmplhc2gupfne2em5qfn0joy"}'
```

## Session Key

After logging in, you receive a session key:

```json
{"status": "success", "message": "authenticated user: ", "session": "575d80cf44c0aba5491645a6818589c6"}
```

Include this session key in all subsequent requests.

## Submitting a URL

**API URL:** `http://nova.astrometry.net/api/url_upload`

### Arguments

| Parameter              | Type    | Description                                  |
|------------------------|---------|----------------------------------------------|
| `session`              | string  | Your session key                             |
| `url`                  | string  | URL to submit                                |
| `allow_commercial_use` | string  | "d", "y", or "n" (licensing)                 |
| `allow_modifications`  | string  | "d", "y", "n", or "sa" (licensing)           |
| `publicly_visible`     | string  | "y" or "n"                                   |
| `scale_units`          | string  | "degwidth", "arcminwidth", or "arcsecperpix" |
| `scale_type`           | string  | "ul" or "ev"                                 |
| `scale_lower`          | float   | Lower scale bound                            |
| `scale_upper`          | float   | Upper scale bound                            |
| `scale_est`            | float   | Scale estimate                               |
| `scale_err`            | float   | Scale error                                  |
| `center_ra`            | float   | Center RA in degrees                         |
| `center_dec`           | float   | Center Dec in degrees                        |
| `radius`               | float   | Radius in degrees                            |
| `downsample_factor`    | float   | Value > 1                                    |
| `tweak_order`          | int     | Tweak order                                  |
| `use_sextractor`       | boolean | Use SExtractor                               |
| `crpix_center`         | boolean | CRPIX center                                 |
| `parity`               | int     | 0, 1, or 2                                   |
| `image_width`          | int     | Image width                                  |
| `image_height`         | int     | Image height                                 |
| `positional_error`     | float   | Positional error                             |

### Example Request

```json
{
  "session": "####",
  "url": "http://apod.nasa.gov/apod/image/1206/ldn673s_block1123.jpg",
  "scale_units": "degwidth",
  "scale_lower": 0.5,
  "scale_upper": 1.0,
  "center_ra": 290,
  "center_dec": 11,
  "radius": 2.0
}
```

### Example Response

```json
{"status": "success", "subid": 16714, "hash": "6024b45a16bfb5af7a73735cbabdf2b462c11214"}
```

## Submitting a File

**API URL:** `http://nova.astrometry.net/api/upload`

Format as `multipart/form-data`:

1. First part: text field `request-json` with JSON
2. Second part: file data with MIME type `octet-stream`

### Example

```
--===============2521702492343980833==
Content-Type: text/plain
MIME-Version: 1.0
Content-disposition: form-data; name="request-json"

{"publicly_visible": "y", "allow_modifications": "d", "session": "XXXXXX", "allow_commercial_use": "d"}
--===============2521702492343980833==
Content-Type: application/octet-stream
MIME-Version: 1.0
Content-disposition: form-data; name="file"; filename="myfile.txt"

Hello World

--===============2521702492343980833==--
```

## Getting Submission Status

**API URL:** `http://nova.astrometry.net/api/submissions/SUBID`

### Example Response

```json
{
  "processing_started": "2016-03-29 11:02:11.967627",
  "job_calibrations": [[1493115, 785516]],
  "jobs": [1493115],
  "processing_finished": "2016-03-29 11:02:13.010625",
  "user": 1,
  "user_images": [1051223]
}
```

## Getting Job Status

**API URL:** `http://nova.astrometry.net/api/jobs/JOBID`

### Example Response

```json
{"status": "success"}
```

## Getting Job Results

### Calibration

**API URL:** `http://nova.astrometry.net/api/jobs/JOBID/calibration/`

```json
{
  "parity": 1.0,
  "orientation": 105.74942079091929,
  "pixscale": 1.0906710701159739,
  "radius": 0.8106715896625917,
  "ra": 169.96633791366915,
  "dec": 13.221011585315143
}
```

### Tagged Objects

**API URLs:**

- `http://nova.astrometry.net/api/jobs/JOBID/tags/`
- `http://nova.astrometry.net/api/jobs/JOBID/machine_tags/`

```json
{"tags": ["NGC 3628", "M 66", "NGC 3627", "M 65", "NGC 3623"]}
```

### Known Objects in Field

**API URL:** `http://nova.astrometry.net/api/jobs/JOBID/objects_in_field/`

```json
{"objects_in_field": ["NGC 3628", "M 66", "NGC 3627", "M 65", "NGC 3623"]}
```

### Objects with Coordinates (Annotations)

**API URL:** `http://nova.astrometry.net/api/jobs/JOBID/annotations/`

```json
{
  "annotations": [
    {
      "radius": 0.0,
      "type": "ic",
      "names": ["IC 2728"],
      "pixelx": 1604.1727638846828,
      "pixely": 1344.045125738614
    },
    {
      "radius": 0.0,
      "type": "hd",
      "names": ["HD 98388"],
      "pixelx": 1930.2719762446786,
      "pixely": 625.1110603737037
    }
  ]
}
```

### Full Job Info

**API URL:** `http://nova.astrometry.net/api/jobs/JOBID/info/`

```json
{
  "status": "success",
  "machine_tags": ["NGC 3628", "M 66", "NGC 3627", "M 65", "NGC 3623"],
  "calibration": {
    "parity": 1.0,
    "orientation": 105.74942079091929,
    "pixscale": 1.0906710701159739,
    "radius": 0.8106715896625917,
    "ra": 169.96633791366915,
    "dec": 13.221011585315143
  },
  "tags": ["NGC 3628", "M 66", "NGC 3627", "M 65", "NGC 3623"],
  "original_filename": "Leo Triplet-1.jpg",
  "objects_in_field": ["NGC 3628", "M 66", "NGC 3627", "M 65", "NGC 3623"]
}
```

## Result Files

| File Type         | URL                                                         |
|-------------------|-------------------------------------------------------------|
| WCS file          | `http://nova.astrometry.net/wcs_file/JOBID`                 |
| New FITS file     | `http://nova.astrometry.net/new_fits_file/JOBID`            |
| RDLS file         | `http://nova.astrometry.net/rdls_file/JOBID`                |
| AXY file          | `http://nova.astrometry.net/axy_file/JOBID`                 |
| Correlation file  | `http://nova.astrometry.net/corr_file/JOBID`                |
| Annotated display | `http://nova.astrometry.net/annotated_display/JOBID`        |
| Red/green image   | `http://nova.astrometry.net/red_green_image_display/JOBID`  |
| Extraction image  | `http://nova.astrometry.net/extraction_image_display/JOBID` |
