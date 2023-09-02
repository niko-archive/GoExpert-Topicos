import 'package:api_consumer/view/home.dart';
import 'package:api_consumer/view/vehicle.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

final GoRouter routes = GoRouter(
  routes: <RouteBase>[
    GoRoute(
      path: '/',
      builder: (BuildContext context, GoRouterState state) {
        return const Home();
      },
    ),
    GoRoute(
      path: '/vehicle',
      builder: (BuildContext context, GoRouterState state) {
        return const VehicleView();
      },
    ),
  ],
);
