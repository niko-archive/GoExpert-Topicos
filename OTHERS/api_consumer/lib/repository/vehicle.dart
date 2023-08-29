// ignore_for_file: avoid_print

import 'package:api_consumer/dto/vehicle.dart';
import 'package:uno/uno.dart';

class RepoVehicle {
  Uno uno = Uno();

  Future<List<VehicleDTO>> getPagination(int page, int limit) async {
    Uri uri = Uri(
      scheme: 'http',
      host: '192.168.3.27',
      port: 8001,
      path: '/vehicles/all',
      queryParameters: {
        'page': '$page',
        'limit': '$limit',
      },
    );

    final api = uri.toString();
    Response response = await uno.get(api);
    if (response.status != 200) {
      print('Error: ${response.status}');
    }

    List<VehicleDTO> vehicles = [];
    if (response.data == null) {
      return vehicles;
    }
    response.data.forEach((vehicle) {
      vehicles.add(VehicleDTO.fromJson(vehicle));
    });

    return vehicles;
  }
}
