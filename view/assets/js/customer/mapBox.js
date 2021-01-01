function getTheMap(placeList) {
    mapboxgl.accessToken = 'pk.eyJ1IjoiZXNwb2lyLWRpdGUiLCJhIjoiY2tmaTZsdDBiMGpubzJzcDh2aXdodXlnayJ9.e6tqu6hLu5xlaSk84ERd9g';
    var map = new mapboxgl.Map({
        container: 'mapz',
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [18.446046,-33.986392],
        zoom: 11.15
    });

    map.on('load', function () {
        //console.log("returns of getPlaceData(placeList: ",getPlaceData(placeList));
        map.addSource('places', {
            'type': 'geojson',
            'data': {
                'type': 'FeatureCollection',
                'features': getPlaceData(placeList)
            }
        });
// Add a layer showing the places.
        map.addLayer({
            'id': 'places',
            'type': 'symbol',
            'source': 'places',
            'layout': {
                'icon-image': '{icon}-15',
                'icon-allow-overlap': truelo
            }
        });

// When a click event occurs on a feature in the places layer, open a popup at the
// location of the feature, with description HTML from its properties.
        map.on('click', 'places', function (e) {
            var coordinates = e.features[0].geometry.coordinates.slice();
            var description = e.features[0].properties.description;

// Ensure that if the map is zoomed out such that multiple
// copies of the feature are visible, the popup appears
// over the copy being pointed to.
            while (Math.abs(e.lngLat.lng - coordinates[0]) > 180) {
                coordinates[0] += e.lngLat.lng > coordinates[0] ? 360 : -360;
            }

            new mapboxgl.Popup()
                .setLngLat(coordinates)
                .setHTML(description)
                .addTo(map);
        });

// Change the cursor to a pointer when the mouse is over the places layer.
        map.on('mouseenter', 'places', function () {
            map.getCanvas().style.cursor = 'pointer';
        });

// Change it back to a pointer when it leaves.
        map.on('mouseleave', 'places', function () {
            map.getCanvas().style.cursor = '';
        });
    });
}














function getPlaceData(PlaceList) {
    //console.log(PlaceList);
    var myPlaceList = [];
    for (let placeList of PlaceList) {
        console.log(placeList);
        var logitude = placeList.longitude
        var  trimedLongitude = logitude.trim();
         myPlaceList.push( {
            'type': 'Feature',
            'properties': {
                'description':
                    '<strong>placeList.title</strong><p><a href="http://www.mtpleasantdc.com/makeitmtpleasant" target="_blank" title="Opens in a new window">Make it Mount Pleasant</a>placeList.description</p>',
                'icon': ''
            },
            'geometry': {
                'type': 'Point',
                'coordinates': [trimedLongitude, placeList.latitude]
            }
        });
    }
    return myPlaceList;
}
