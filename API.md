# API Documentation

Complete API reference for the Face Detection App backend.

## Base URL

```
http://localhost:8080
```

## Authentication

The API currently does not require authentication. All endpoints are public.

## Response Format

All responses are returned in JSON format.

### Success Response Format

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {}
}
```

### Error Response Format

```json
{
  "success": false,
  "message": "Error description"
}
```

## Endpoints

### 1. Health Check

Verifies that the backend API is running and healthy.

**Endpoint:** `GET /api/health`

**Description:** Returns the current health status of the API service.

**Request:**
```bash
curl http://localhost:8080/api/health
```

**Response:** `200 OK`
```json
{
  "status": "ok"
}
```

**Use Cases:**
- Kubernetes health probes
- Docker container health checks
- Load balancer monitoring
- Application startup verification

---

### 2. Detect Faces

Analyzes an image and detects faces using OpenCV Haar Cascade.

**Endpoint:** `POST /api/detect-face`

**Description:** 
- Accepts a base64-encoded image
- Detects faces using Haar Cascade classifier
- Returns the image with detected faces highlighted
- Includes count of detected faces

**Request:**

**Headers:**
```
Content-Type: application/json
```

**Body:**
```json
{
  "imageData": "base64_encoded_image_string"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/detect-face \
  -H "Content-Type: application/json" \
  -d '{
    "imageData": "/9j/4AAQSkZJRgABAQEAYABgAAD/..."
  }'
```

**Response:** `200 OK` (Success)
```json
{
  "success": true,
  "imageData": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD/...",
  "faceCount": 2,
  "message": "Detected 2 face(s)"
}
```

**Response:** `400 Bad Request` (Invalid Input)
```json
{
  "success": false,
  "message": "Invalid request: Field 'imageData' is required"
}
```

**Response:** `400 Bad Request` (Invalid Base64)
```json
{
  "success": false,
  "message": "Failed to decode image: illegal base64 data at input byte X"
}
```

**Response:** `500 Internal Server Error` (Detection Failed)
```json
{
  "success": false,
  "message": "Face detection failed: failed to load face classifier"
}
```

**Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `imageData` | string | Yes | Base64-encoded image data (JPEG, PNG, etc.) |

**Response Fields:**
| Field | Type | Description |
|-------|------|-------------|
| `success` | boolean | Whether the operation was successful |
| `imageData` | string | Base64-encoded result image with faces highlighted (includes `data:image/jpeg;base64,` prefix) |
| `faceCount` | integer | Number of faces detected in the image |
| `message` | string | Human-readable status message |

**Status Codes:**
| Code | Meaning |
|------|---------|
| 200 | Successful face detection |
| 400 | Bad request (missing/invalid imageData) |
| 500 | Server error (detection failed) |

**Image Requirements:**
- **Format:** JPEG, PNG, or other formats supported by Go image library
- **Size:** Recommended maximum 5MB
- **Dimensions:** No hard limits, but larger images take longer to process
- **Encoding:** Must be valid Base64

**Face Detection Notes:**
- Uses OpenCV Haar Cascade Classifier for detection
- Detected faces are highlighted with green rectangles
- Cascade file: `haarcascade_frontalface_default.xml`
- Detection parameters can be tuned in code if needed

**Example Usage in JavaScript:**

```javascript
async function detectFaces(imageFile) {
  // Read file as Base64
  const reader = new FileReader();
  reader.readAsDataURL(imageFile);
  
  reader.onload = async (event) => {
    // Remove "data:image/jpeg;base64," prefix
    const base64String = event.target.result.split(',')[1];
    
    // Call API
    const response = await fetch('http://localhost:8080/api/detect-face', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        imageData: base64String
      })
    });
    
    const data = await response.json();
    
    if (data.success) {
      console.log(`Detected ${data.faceCount} face(s)`);
      // Display result image
      const img = new Image();
      img.src = data.imageData;
      document.body.appendChild(img);
    } else {
      console.error(data.message);
    }
  };
}
```

**Example Usage in Python:**

```python
import requests
import base64

def detect_faces(image_path):
    # Read image file and encode to base64
    with open(image_path, 'rb') as image_file:
        base64_string = base64.b64encode(image_file.read()).decode('utf-8')
    
    # Call API
    response = requests.post(
        'http://localhost:8080/api/detect-face',
        json={'imageData': base64_string},
        headers={'Content-Type': 'application/json'}
    )
    
    data = response.json()
    
    if data['success']:
        print(f"Detected {data['faceCount']} face(s)")
        # Process result image if needed
    else:
        print(f"Error: {data['message']}")
    
    return data
```

**Example Usage in cURL:**

```bash
# Convert image to base64
base64_image=$(base64 < /path/to/image.jpg | tr -d '\n')

# Send to API
curl -X POST http://localhost:8080/api/detect-face \
  -H "Content-Type: application/json" \
  -d "{\"imageData\": \"$base64_image\"}" \
  -o result.json

# View result
cat result.json
```

---

## Error Handling

### Common Error Scenarios

**Missing Required Field:**
```json
{
  "success": false,
  "message": "Invalid request: Field 'imageData' is required"
}
```

**Invalid Base64 Encoding:**
```json
{
  "success": false,
  "message": "Failed to decode image: illegal base64 data at input byte 42"
}
```

**Unsupported Image Format:**
```json
{
  "success": false,
  "message": "Failed to decode image format: unknown format"
}
```

**Classifier Not Found:**
```json
{
  "success": false,
  "message": "Face detection failed: failed to load face classifier"
}
```

### Error Recovery

1. **Verify Base64 Encoding:**
   ```javascript
   // Ensure imageData is proper base64 without prefix
   const cleanBase64 = imageData.replace(/^data:.*?;base64,/, '');
   ```

2. **Check Image Format:**
   ```javascript
   // Ensure image is JPEG or PNG
   const validFormats = ['image/jpeg', 'image/png'];
   if (!validFormats.includes(imageFile.type)) {
     throw new Error('Unsupported image format');
   }
   ```

3. **Validate Image Size:**
   ```javascript
   // Limit image size to 5MB
   const maxSize = 5 * 1024 * 1024; // 5MB
   if (imageFile.size > maxSize) {
     throw new Error('Image too large');
   }
   ```

---

## Rate Limiting

Currently, there are no rate limits on the API endpoints. For production deployments, consider implementing:

- Request rate limiting
- IP-based throttling
- API keys for access control

---

## CORS Configuration

The API has CORS (Cross-Origin Resource Sharing) enabled with the following configuration:

```
Allowed Origins: *
Allowed Methods: GET, POST, OPTIONS
Allowed Headers: Content-Type
```

This allows the React frontend to communicate with the backend without CORS issues.

---

## Performance Considerations

### Image Processing Time

Face detection performance depends on:
- Image resolution
- Number of faces in image
- System hardware
- OpenCV optimization level

**Typical Processing Times:**
- VGA (640x480): 50-100ms
- Full HD (1920x1080): 200-500ms
- 4K (3840x2160): 1000-3000ms

### Optimization Tips

1. **Resize Large Images:**
   ```javascript
   function resizeImage(file, maxWidth = 800) {
     const canvas = document.createElement('canvas');
     const ctx = canvas.getContext('2d');
     const img = new Image();
     
     img.onload = () => {
       const ratio = maxWidth / img.width;
       canvas.width = maxWidth;
       canvas.height = img.height * ratio;
       ctx.drawImage(img, 0, 0, canvas.width, canvas.height);
       return canvas.toDataURL('image/jpeg', 0.8);
     };
     img.src = URL.createObjectURL(file);
   }
   ```

2. **Compress Images:**
   - Use JPEG with quality 75-85%
   - Reduce resolution to necessary size
   - Strip metadata

3. **Batch Processing:**
   - For multiple images, consider async processing
   - Implement retry logic for failed requests

---

## API Versioning

Current API Version: **v1** (implicit)

Future versions may be prefixed as `/api/v2/`, maintaining backward compatibility.

---

## Technology Stack

- **Framework:** Gin Web Framework
- **Language:** Go 1.21
- **Computer Vision:** OpenCV with GoCV bindings
- **Format:** JSON
- **CORS:** Gin CORS middleware

---

## Related Documentation

- [README.md](README.md) - Project overview
- [INSTALLATION.md](INSTALLATION.md) - Setup instructions
- [DEVELOPMENT.md](DEVELOPMENT.md) - Development guide
- [DOCKER.md](DOCKER.md) - Docker deployment
