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
        // leading: backToHome(context),
        actions: [
          IconButton(
            onPressed: () {
              context.go(Home.routeName);
            },
            icon: const Icon(Icons.home),
          ),
        ],
      ),
      body: const Center(
        child: Column(
          children: [
            ListOfVehicles(),
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

  bool isLoading = false;
  bool firstLoad = true;
  RepoVehicle repo = RepoVehicle();
  List<VehicleDTO> vehicles = [];

  final ScrollController scroll = ScrollController();

  @override
  void initState() {
    super.initState();
    firstPopulated();

    scroll.addListener(() async {
      var maxScroll = scroll.position.maxScrollExtent;
      var currentScroll = scroll.position.pixels;

      if (currentScroll == maxScroll) {
        isLoading = true;
        repo.getPagination(page, 20).then((values) {
          setState(() {
            vehicles.addAll(values);
          });
        });
        page++;
        isLoading = false;
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
      future: () async {
        return !isLoading;
      }(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return Expanded(
            child: Scrollbar(
              controller: scroll,
              child: ListView.builder(
                physics: const ClampingScrollPhysics(),
                controller: scroll,
                itemCount: vehicles.length,
                itemBuilder: (context, index) {
                  return ListTile(
                    title: Text(vehicles[index].model),
                    subtitle: Text(vehicles[index].brand),
                    leading: CircleAvatar(
                      child: Text(vehicles[index].fuel.substring(0, 1)),
                    ),
                  );
                },
              ),
            ),
          );
        } else {
          return const Center(
            child: CircularProgressIndicator(),
          );
        }
      },
    );
  }

  void firstPopulated() {
    if (firstLoad) {
      isLoading = true;
      repo.getPagination(page, 15).then((values) {
        setState(() {
          vehicles.addAll(values);
          page++;
        });
      });
      isLoading = false;
      firstLoad = false;
    }
  }
}
