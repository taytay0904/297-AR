using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;

public class DroneAnimate : MonoBehaviour
{
	public Animator anim;
    public Animator animDL;
    // Start is called before the first frame update

    public IDictionary modelList = new Dictionary<string, Animator>();

    void Start()
    {
        Debug.Log("Start");
        // string url = "localhost:8080/model";
        // WWW www = new WWW(url); 
        StartCoroutine(WaitForRes());       
    }

    IEnumerator WaitForRes(){
        while (!QRdetector.load){
            yield return null; 
        }
        Debug.Log("Start Downloading");
        Debug.Log(QRdetector.modelAddress);
        // IDictionary modelAddress = (IDictionary)QRdetector.modelAddress;
        foreach(KeyValuePair<string, string> entry in QRdetector.modelAddress){
            Debug.Log("Address: "+entry.Value);
            Debug.Log("Model name: "+entry.Key);
            WWW www = new WWW(entry.Value); 
            yield return www;
            try
            {           
                Debug.Log(www.assetBundle);
                AssetBundle bundel = www.assetBundle;
                if (bundel){                    
                    Debug.Log("Download finish");
                    animDL = ((GameObject)bundel.LoadAsset(entry.Key)).GetComponent<Animator> ();
                    Debug.Log("read finish");
                    Instantiate(animDL);
                    modelList.Add(entry.Key,animDL);
                    Debug.Log(animDL.GetCurrentAnimatorStateInfo(0));
                }
                else{
                    Debug.Log("Error");
                    Debug.Log(www.error);
                }
            }
            catch (Exception e)
            {
                Debug.LogError(e);
                Debug.LogError(e.Message);
            }
        }      
        Debug.Log("Finish Downloading");
        
    }

    // Update is called once per frame
    void Update()
    {
        if(Input.GetKeyDown("w"))
        {
            anim.Play("PA_DroneLeanFoward_Clip");
        }else if(Input.GetKeyDown("s") )
        {
            anim.Play("PA_DroneBackwards_Clip");
        }else if(Input.GetKeyDown("a") )
        {
            anim.Play("PA_DroneMoveLeft_Clip");
        }else if(Input.GetKeyDown("d") )
        {
            anim.Play("PA_DroneMoveRight_Clip");
        }else if(Input.GetKeyDown("p") )
        {
            anim.Play("PA_DroneExplode_Clip");
        }



        // for (int i = 0; i < Input.touchCount; ++i){
        //     if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("none"))
        //     {
        //         anim.Play("PA_DroneLeanFoward_Clip");
        //     }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneLeanFoward_Clip"))
        //     {
        //         anim.Play("PA_DroneBackwards_Clip");
        //     }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneBackwards_Clip"))
        //     {
        //         anim.Play("PA_DroneMoveLeft_Clip");
        //     }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneMoveLeft_Clip"))
        //     {
        //         anim.Play("PA_DroneMoveRight_Clip");
        //     }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneMoveRight_Clip"))
        //     {
        //         anim.Play("PA_DroneExplode_Clip");
        //     }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneExplode_Clip"))
        //     {
        //         anim.Play("none");
        //     }
        // }



    }
}
