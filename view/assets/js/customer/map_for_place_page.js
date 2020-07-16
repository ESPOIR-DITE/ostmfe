

function initMap(PlaceList) {
    // console.log(PlaceList);
    // Object.keys(PlaceList).forEach((key, index) => {
    //     console.log(key ,PlaceList[key]);
    // });
    // for (var key in PlaceList) {
    //     if (PlaceList.hasOwnProperty(key)) {
    //         console.log(key + " -> " + PlaceList[key]);
    //     }
    // }
    // var myLatLng = {lat: -25.363, lng: 131.044};
    //
    // var map = new google.maps.Map(document.getElementById('map'), {
    //     zoom: 4,
    //     center: myLatLng
    // });
    //
    // var marker = new google.maps.Marker({
    //     position: myLatLng,
    //     map: map,
    //     title: 'Hello World!'
    // });
}
// <script>
// var markers = null;
// var infoWindowContent = null
// var myMarkerList = new Set();
// var myInfoWindowList = new Set();
// function initMap() {
//     var PlaceList ={{.Places}};
//     Object.keys(PlaceList).forEach((key,index)=>{
//         console.log(key,PlaceList[key])
//         myMarkerList.add([PlaceList[key].title,PlaceList[key].latitude, PlaceList[key].longitude]);
//         myInfoWindowList.add(['<div class="info_content">' +
//         '<h3>Brooklyn Museum</h3>' +
//         '<p>The Brooklyn Museum is an art museum located in the New York City borough of Brooklyn.</p>' + '</div>'],);
//     })
//     console.log('myMarkerList :',myMarkerList);
//     var map;
//     var bounds = new google.maps.LatLngBounds();
//     var mapOptions = {
//         mapTypeId: 'roadmap'
//     };
//
//
//     // Display a map on the web page
//     map = new google.maps.Map(document.getElementById("mapCanvas"), mapOptions);
//     map.setTilt(50);
//
//     // Multiple markers location, latitude, and longitude
//     markers = myMarkerList;
//
//     // Info window content
//     infoWindowContent = myInfoWindowList;
//
//     // Add multiple markers to map
//     var infoWindow = new google.maps.InfoWindow(), marker, i;
//
//     // Place each marker on the map
//     for( i = 0; i < markers.length; i++ ) {
//         var position = new google.maps.LatLng(markers[i][1], markers[i][2]);
//         bounds.extend(position);
//         marker = new google.maps.Marker({
//             position: position,
//             map: map,
//             title: markers[i][0]
//         });
//
//         // Add info window to marker
//         google.maps.event.addListener(marker, 'click', (function(marker, i) {
//             return function() {
//                 infoWindow.setContent(infoWindowContent[i][0]);
//                 infoWindow.open(map, marker);
//             }
//         })(marker, i));
//
//         // Center the map to fit all markers on the screen
//         map.fitBounds(bounds);
//     }
//
//     // Set zoom level
//     var boundsListener = google.maps.event.addListener((map), 'bounds_changed', function(event) {
//         this.setZoom(14);
//         google.maps.event.removeListener(boundsListener);
//     });
//
// }
// // Load initialize function
// google.maps.event.addDomListener(window, 'load', initMap);
// </script>
