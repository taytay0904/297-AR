using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ObjectSpawner : MonoBehaviour
{
    public GameObject objectTOSpawn;
    private PlacementIndicator PlacementIndicator;
    // Start is called before the first frame update
    void Start()
    {
        PlacementIndicator = FindObjectOfType<PlacementIndicator>();
    }

    // Update is called once per frame
    void Update()
    {
        if(Input.touchCount > 0 && Input.touches[0].phase == TouchPhase.Began){
            GameObject obj = Instantiate(objectTOSpawn,PlacementIndicator.transform.position, PlacementIndicator.transform.rotation);
        }
    }
}
