function initMap(listener) {

    var myLatlng = {lat: -34.192697, lng: 18.427731};

    var latitudeField = document.getElementById('exampleRepeatPassword');

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