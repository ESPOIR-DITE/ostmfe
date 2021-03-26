mapboxgl.accessToken = 'pk.eyJ1IjoiZXNwb2lyLWRpdGUiLCJhIjoiY2tmaTZsdDBiMGpubzJzcDh2aXdodXlnayJ9.e6tqu6hLu5xlaSk84ERd9g';
var geojson = {
    'type': 'FeatureCollection',
    'features': [
        {
            'type': 'Feature',
            'properties': {
                'message': 'Foo',
                'iconSize': [60, 60]
            },
            'geometry': {
                'type': 'Point',
                'coordinates': [-66.324462890625, -16.024695711685304]
            }
        },
        {
            'type': 'Feature',
            'properties': {
                'message': 'Bar',
                'iconSize': [50, 50]
            },
            'geometry': {
                'type': 'Point',
                'coordinates': [-61.2158203125, -15.97189158092897]
            }
        },
        {
            'type': 'Feature',
            'properties': {
                'message': 'Baz',
                'iconSize': [40, 40]
            },
            'geometry': {
                'type': 'Point',
                'coordinates': [-63.29223632812499, -18.28151823530889]
            }
        }
    ]
};

var myPlaces = getPlaceData();
console.log("myplace: "+myPlaces)

var map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [-65.017, -16.457],
    zoom: 5
});

// add markers to map
geojson.features.forEach(function (marker) {
// create a DOM element for the marker
    var el = document.createElement('div');
    el.className = 'marker';
    // el.style.backgroundImage = '../../assets/img/marker-15.svg';
    el.style.width = marker.properties.iconSize[0] + 'px';
    el.style.height = marker.properties.iconSize[1] + 'px';

    el.addEventListener('click', function () {
        // console.log("voila");
        $('#title').setAttribute('text','voila');
        $title.set('test','voila');
        window.alert("marker.properties.message");
    });

    el.addEventListener('mousemove', function () {
        el.setAttribute('title','voila!!');
    });
// add marker to map
    var marker = new mapboxgl.Marker(el)
        .setLngLat(marker.geometry.coordinates)
        .addTo(map);

});

function getPlaceData(PlaceList) {
    //console.log(PlaceList);
    var myPlaceList = [];
    for (let placeList of PlaceList) {
        console.log(placeList);
        var logitude = placeList.longitude
        var  trimedLongitude = logitude.trim();
        myPlaceList.push( {
            'type': 'FeatureCollection',
            'features': {
                'type': 'Feature',
                'properties': {
                    'message':'<strong>placeList.title</strong><p><a href="http://www.mtpleasantdc.com/makeitmtpleasant" target="_blank" title="Opens in a new window">Make it Mount Pleasant</a>placeList.description</p>',
                    'iconSize': [60, 60]
                },
                'geometry': {
                    'type': 'Point',
                    'coordinates': [trimedLongitude, placeList.latitude]
                }
            }
        });
    }
    return myPlaceList;
}
