using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;

using Vuforia;

using System.Threading;

using ZXing;
using ZXing.QrCode;
using ZXing.Common;

[AddComponentMenu("System/VuforiaScanner")]
public class QRdetector : MonoBehaviour
{
    private bool cameraInitialized;

    private BarcodeReader barCodeReader;

    private PIXEL_FORMAT imageFormat = PIXEL_FORMAT.GRAYSCALE;

    private String address = null;

    public static bool load = false;

    public static List<KeyValuePair<string, string>> modelAddress = new List<KeyValuePair<string, string>>();
       
    // Start is called before the first frame update
    void Start()
    {   
        // address = "http://ec2-13-52-250-41.us-west-1.compute.amazonaws.com:8080/";
        // modelAddress.Add(new KeyValuePair<string, string>("pa_warrior", address + "pa_warrior"));
        // modelAddress.Add(new KeyValuePair<string, string>("pa_drone", address + "pa_drone"));
        // load = true;

        barCodeReader = new BarcodeReader();
        StartCoroutine(InitializeCamera());
    }

    private IEnumerator InitializeCamera()
    {
        // Waiting a little seem to avoid the Vuforia's crashes.
        yield return new WaitForSeconds(1.25f);

        var isFrameFormatSet = CameraDevice.Instance.SetFrameFormat(imageFormat, true);
        Debug.Log(String.Format("FormatSet : {0}", isFrameFormatSet));

        // Force autofocus.
        var isAutoFocus = CameraDevice.Instance.SetFocusMode(CameraDevice.FocusMode.FOCUS_MODE_CONTINUOUSAUTO);
        if (!isAutoFocus)
        {
            CameraDevice.Instance.SetFocusMode(CameraDevice.FocusMode.FOCUS_MODE_NORMAL);
        }
        Debug.Log(String.Format("AutoFocus : {0}", isAutoFocus));
        cameraInitialized = true;
    }

    // Update is called once per frame
    void Update()
    {   
        // Debug.Log(String.Format("cameraInitialized : {0}", cameraInitialized));
        
        if (cameraInitialized && address==null)
        {
            try
            {
                var cameraFeed = CameraDevice.Instance.GetCameraImage(imageFormat);
                if (cameraFeed == null)
                {
                    Debug.Log("No Image");
                    return;
                }
                var data = barCodeReader.Decode(cameraFeed.Pixels, cameraFeed.BufferWidth, cameraFeed.BufferHeight, RGBLuminanceSource.BitmapFormat.RGB24);
                if (data != null)
                {
                    // QRCode detected.
                    Debug.Log(data.Text);
                    address = "http://ec2-13-52-250-41.us-west-1.compute.amazonaws.com:8080/";
                    modelAddress.Add(new KeyValuePair<string, string>("pa_warrior", address + "pa_warrior"));
                    modelAddress.Add(new KeyValuePair<string, string>("pa_drone", address + "pa_drone"));
                    load = true;
                }
                else
                {
                    Debug.Log("No QR code detected !");
                }
            }
            catch (Exception e)
            {
                Debug.LogError(e);
                Debug.LogError(e.Message);
            }
        }
    }
}
