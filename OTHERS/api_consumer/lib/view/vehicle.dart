import 'dart:async';

import 'package:api_consumer/store/vehicles.dart';
import 'package:api_consumer/view/home.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class VehicleView extends StatelessWidget {
  VehicleView({super.key});

  final VehiclesStore store = VehiclesStore();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: GestureDetector(
          child: Text('Vehicle: ${store.total}'),
          onTap: () {
            Timer.periodic(const Duration(seconds: 1), (timer) {
              (context as Element).markNeedsBuild();
            });
          },
        ),
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
  VehiclesStore store = VehiclesStore();

  @override
  void initState() {
    super.initState();
    firstPopulated();

    store.scroll.addListener(() async {
      var maxScroll = store.scroll.position.maxScrollExtent;
      var currentScroll = store.scroll.position.pixels;

      if (currentScroll == maxScroll) {
        setState(() {});
        store.isLoading = true;

        store.repo.getPagination(store.page, 50).then((values) {
          store.vehicles.addAll(values);
        });
        store.isLoading = false;
        store.page++;
        setState(() {});
      }
    });
  }

  @override
  void dispose() {
    store.scroll.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: () async {}(),
      builder: (context, snapshot) {
        if (!store.isLoading) {
          return Expanded(
            flex: 1,
            child: Scrollbar(
              thickness: 10,
              controller: store.scroll,
              child: ListView.builder(
                physics: const ClampingScrollPhysics(),
                controller: store.scroll,
                itemCount: store.vehicles.length,
                itemBuilder: (context, index) {
                  return ListTile(
                    title: Text(store.vehicles[index].model),
                    subtitle: Text(store.vehicles[index].brand),
                    leading: CircleAvatar(
                      child: Text(store.vehicles[index].fuel.substring(0, 1)),
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
    if (store.firstLoad) {
      store.repo.getPagination(store.page, 50).then((values) {
        setState(() {
          store.vehicles.addAll(values);
          store.page++;
        });
      });

      store.firstLoad = false;
    }
  }
}
