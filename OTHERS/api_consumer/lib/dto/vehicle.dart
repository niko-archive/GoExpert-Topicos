class VehicleDTO {
  final String id;
  final int value;
  final String brand;
  final String model;
  final String modelYear;
  final String fuel;
  final String fipeCode;
  final String referenceMonth;
  final String vehicleType;

  VehicleDTO({
    required this.id,
    required this.value,
    required this.brand,
    required this.model,
    required this.modelYear,
    required this.fuel,
    required this.fipeCode,
    required this.referenceMonth,
    required this.vehicleType,
  });

  factory VehicleDTO.fromJson(Map<String, dynamic> json) {
    return VehicleDTO(
      id: json['id'],
      value: json['value'],
      brand: json['brand'],
      model: json['model'],
      modelYear: json['model_year'],
      fuel: json['fuel'],
      fipeCode: json['fipe_code'],
      referenceMonth: json['reference_month'],
      vehicleType: json['vehicle_type'],
    );
  }
}
