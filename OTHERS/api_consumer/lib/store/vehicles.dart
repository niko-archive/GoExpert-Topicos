import 'package:api_consumer/dto/vehicle.dart';
import 'package:api_consumer/repository/vehicle.dart';
import 'package:flutter/material.dart';

class VehiclesStore {
  // Private constructor
  VehiclesStore._();
  // Singleton pattern
  static final VehiclesStore _instance = VehiclesStore._();

  factory VehiclesStore() => _instance;

  int page = 1;
  bool isLoading = false;
  bool firstLoad = true;
  RepoVehicle repo = RepoVehicle();
  List<VehicleDTO> vehicles = [];

  final ScrollController scroll = ScrollController();
  get total => vehicles.length;
}
