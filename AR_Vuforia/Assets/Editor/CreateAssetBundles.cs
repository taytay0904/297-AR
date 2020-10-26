using System.Collections;
using System.Collections.Generic;
using UnityEditor;
using UnityEngine;


public class CreateAssetBundles : MonoBehaviour
{
    [MenuItem ("Assets/Build AssetBundles-Windows")]
    static void BuildAllAssetBundles(){
        BuildPipeline.BuildAssetBundles ("Assets/AssetBundles",BuildAssetBundleOptions.None, BuildTarget.StandaloneWindows);
    }
    
}
