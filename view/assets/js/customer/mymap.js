function initMap(listener) {

    var myLatlng = {lat: -34.192697, lng: 18.427731};

    var latitudeField = document.getElementById('exampleRepeatPassword');

    var map = new google.maps.Map(document.getElementById('mapx'), {zoom: 15, center: myLatlng});

    // Create the initial InfoWindow.
    var infoWindow = new google.maps.InfoWindow(
        {content: 'Click the map to get Lat/Lng!', position: myLatlng});
    infoWindow.open(map);

    // Configure the click listener.
    map.addListener('click', function (mapsMouseEvent) {
        //enabling the field

        // Close the current InfoWindow.
        infoWindow.close();

        // Create a new InfoWindow.
        infoWindow = new google.maps.InfoWindow({position: mapsMouseEvent.latLng});
        infoWindow.setContent(mapsMouseEvent.latLng.toString());
        latitudeField.value = mapsMouseEvent.latLng.toString();

        infoWindow.open(map);
    });
}
function mapForPlacePage(listener) {

    var myLatlng = {lat: -34.192697, lng: 18.427731};

    var latitudeField = document.getElementById('latlng');

    $("#exampleRepeatPassword").removeAttr('disabled');
    var map = new google.maps.Map(document.getElementById('map'), {zoom: 15, center: myLatlng});

    // Create the initial InfoWindow.
    var infoWindow = new google.maps.InfoWindow(
        {content: 'Click the map to get Lat/Lng!', position: myLatlng});
        infoWindow.open(map);

    // Configure the click listener.
    map.addListener('click', function (mapsMouseEvent) {
        // Close the current InfoWindow.
        infoWindow.close();

        // Create a new InfoWindow.
        infoWindow = new google.maps.InfoWindow({position: mapsMouseEvent.latLng});
        infoWindow.setContent(mapsMouseEvent.latLng.toString());
        latitudeField.value = mapsMouseEvent.latLng.toString();
        infoWindow.open(map);
    });
}

function initMap() {
    mapboxgl.accessToken = 'pk.eyJ1IjoiZXNwb2lyLWRpdGUiLCJhIjoiY2tmaTZsdDBiMGpubzJzcDh2aXdodXlnayJ9.e6tqu6hLu5xlaSk84ERd9g';
    var map = new mapboxgl.Map({
        container: 'mapx', // container id
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [-74.5, 40], // starting position
        zoom: 9 // starting zoom
    });

    map.on('mousemove', function (e) {
        document.getElementById('latlng').innerHTML =
// e.point is the x, y coordinates of the mousemove event relative
// to the top-left corner of the map
            JSON.stringify(e.point) +
            '<br />' +
            // e.lngLat is the longitude, latitude geographical position of the event
            JSON.stringify(e.lngLat.wrap());
    });
}
