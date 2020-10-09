using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DroneAnimate : MonoBehaviour
{
	public Animator anim;
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        // if(Input.GetKeyDown("w"))
        // {
        //     anim.Play("PA_DroneLeanFoward_Clip");
        // }else if(Input.GetKeyDown("s") )
        // {
        //     anim.Play("PA_DroneBackwards_Clip");
        // }else if(Input.GetKeyDown("a") )
        // {
        //     anim.Play("PA_DroneMoveLeft_Clip");
        // }else if(Input.GetKeyDown("d") )
        // {
        //     anim.Play("PA_DroneMoveRight_Clip");
        // }else if(Input.GetKeyDown("p") )
        // {
        //     anim.Play("PA_DroneExplode_Clip");
        // }



        for (int i = 0; i < Input.touchCount; ++i){
            if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("none"))
            {
                anim.Play("PA_DroneLeanFoward_Clip");
            }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneLeanFoward_Clip"))
            {
                anim.Play("PA_DroneBackwards_Clip");
            }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneBackwards_Clip"))
            {
                anim.Play("PA_DroneMoveLeft_Clip");
            }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneMoveLeft_Clip"))
            {
                anim.Play("PA_DroneMoveRight_Clip");
            }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneMoveRight_Clip"))
            {
                anim.Play("PA_DroneExplode_Clip");
            }else if(Input.GetTouch(i).phase == TouchPhase.Began && anim.GetCurrentAnimatorStateInfo(0).IsName("PA_DroneExplode_Clip"))
            {
                anim.Play("none");
            }
        }



    }
}
