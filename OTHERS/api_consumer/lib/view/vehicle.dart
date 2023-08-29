import 'package:api_consumer/dto/vehicle.dart';
import 'package:api_consumer/repository/vehicle.dart';
import 'package:api_consumer/view/home.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class VehicleView extends StatelessWidget {
  const VehicleView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Vehicle'),
        centerTitle: true,
        backgroundColor: Colors.purple,
        leading: backToHome(context),
      ),
      body: const Center(
        child: Column(
          children: [
            SizedBox(height: 10),
            Expanded(child: ListOfVehicles()),
          ],
        ),
      ),
    );
  }

  IconButton backToHome(BuildContext context) {
    return IconButton(
      onPressed: () {
        context.go(Home.routeName);
      },
      icon: const Icon(Icons.arrow_back),
    );
  }
}

class ListOfVehicles extends StatefulWidget {
  const ListOfVehicles({
    super.key,
  });

  @override
  State<ListOfVehicles> createState() => _ListOfVehiclesState();
}

class _ListOfVehiclesState extends State<ListOfVehicles> {
  int page = 1;

  RepoVehicle repo = RepoVehicle();
  List<VehicleDTO> vehicles = [];

  final ScrollController scroll = ScrollController();

  @override
  void initState() {
    super.initState();
    scroll.addListener(() async {
      var maxScroll = scroll.position.maxScrollExtent;
      var currentScroll = scroll.position.pixels;

      if (currentScroll == maxScroll) {
        page++;
        List<VehicleDTO> vehicles = await repo.getPagination(page, 10);
        addMoreToListDTO(vehicles);
      }
    });
  }

  @override
  void dispose() {
    scroll.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: repo.getPagination(page, 10),
      builder:
          (BuildContext context, AsyncSnapshot<List<VehicleDTO>> snapshot) {
        if (snapshot.hasData) {
          addMoreToListDTO(snapshot.data!);
          return ListView.builder(
            controller: scroll,
            itemCount: vehicles.length,
            shrinkWrap: true,
            itemBuilder: (BuildContext context, int index) {
              // addMoreToList(snapshot, index);
              return ListTile(
                title: Text(vehicles[index].id),
                subtitle: Text(
                    '${vehicles[index].referenceMonth} -- ${vehicles[index].fipeCode} -- ${vehicles[index].vehicleType} -- ${vehicles[index].fuel} -- ${vehicles[index].brand} -- ${vehicles[index].model}'),
              );
            },
          );
        } else {
          return const Center(child: CircularProgressIndicator());
        }
      },
    );
  }

  void addMoreToListDTO(List<VehicleDTO> snapshot) {
    vehicles.addAll(snapshot);
  }
}
