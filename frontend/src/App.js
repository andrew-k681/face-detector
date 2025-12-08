import React, { useRef, useState, useCallback } from 'react';
import './App.css';

function App() {
  const videoRef = useRef(null);
  const canvasRef = useRef(null);
  const [stream, setStream] = useState(null);
  const [capturedImage, setCapturedImage] = useState(null);
  const [detectedImage, setDetectedImage] = useState(null);
  const [faceCount, setFaceCount] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isCameraActive, setIsCameraActive] = useState(false);

  const startCamera = useCallback(async () => {
    try {
      setError(null);
      const mediaStream = await navigator.mediaDevices.getUserMedia({
        video: { 
          width: { ideal: 1280 },
          height: { ideal: 720 },
          facingMode: 'user' 
        }
      });
      
      if (videoRef.current) {
        videoRef.current.srcObject = mediaStream;
        setStream(mediaStream);
        setIsCameraActive(true);
      }
    } catch (err) {
      setError('Failed to access camera: ' + err.message);
      console.error('Camera error:', err);
    }
  }, []);

  const stopCamera = useCallback(() => {
    if (stream) {
      stream.getTracks().forEach(track => track.stop());
      setStream(null);
      setIsCameraActive(false);
    }
    if (videoRef.current) {
      videoRef.current.srcObject = null;
    }
  }, [stream]);

  const capturePhoto = useCallback(() => {
    if (!videoRef.current || !canvasRef.current) return;

    const video = videoRef.current;
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');

    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    context.drawImage(video, 0, 0);

    const imageData = canvas.toDataURL('image/jpeg');
    setCapturedImage(imageData);
  }, []);

  const detectFace = useCallback(async () => {
    if (!capturedImage) {
      setError('Please capture a photo first');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      // Extract base64 data (remove data:image/jpeg;base64, prefix)
      const base64Data = capturedImage.split(',')[1];

      const response = await fetch('/api/detect-face', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          imageData: base64Data,
        }),
      });

      const data = await response.json();

      if (!response.ok || !data.success) {
        throw new Error(data.message || 'Face detection failed');
      }

      setDetectedImage(data.imageData);
      setFaceCount(data.faceCount);
    } catch (err) {
      setError('Face detection error: ' + err.message);
      console.error('Detection error:', err);
    } finally {
      setLoading(false);
    }
  }, [capturedImage]);

  const reset = useCallback(() => {
    setCapturedImage(null);
    setDetectedImage(null);
    setFaceCount(0);
    setError(null);
  }, []);

  return (
    <div className="App">
      <div className="container">
        <h1>Face Detection App</h1>
        
        <div className="controls">
          {!isCameraActive ? (
            <button onClick={startCamera} className="btn btn-primary">
              Start Camera
            </button>
          ) : (
            <>
              <button onClick={stopCamera} className="btn btn-secondary">
                Stop Camera
              </button>
              <button onClick={capturePhoto} className="btn btn-primary" disabled={!isCameraActive}>
                Capture Photo
              </button>
            </>
          )}
          
          {capturedImage && (
            <>
              <button 
                onClick={detectFace} 
                className="btn btn-success" 
                disabled={loading}
              >
                {loading ? 'Detecting...' : 'Detect Face'}
              </button>
              <button onClick={reset} className="btn btn-secondary">
                Reset
              </button>
            </>
          )}
        </div>

        {error && <div className="error-message">{error}</div>}

        <div className="content-grid">
          <div className="camera-section">
            <h2>Camera</h2>
            <div className="video-container">
              <video
                ref={videoRef}
                autoPlay
                playsInline
                muted
                className="video-preview"
              />
              <canvas ref={canvasRef} style={{ display: 'none' }} />
            </div>
          </div>

          <div className="result-section">
            <h2>Captured Photo</h2>
            {capturedImage ? (
              <div className="image-container">
                <img src={capturedImage} alt="Captured" className="result-image" />
              </div>
            ) : (
              <div className="placeholder">No photo captured yet</div>
            )}
          </div>

          {detectedImage && (
            <div className="result-section">
              <h2>Face Detection Result</h2>
              <div className="image-container">
                <img src={detectedImage} alt="Detected faces" className="result-image" />
                <div className="face-count-badge">
                  {faceCount} {faceCount === 1 ? 'Face' : 'Faces'} Detected
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;

