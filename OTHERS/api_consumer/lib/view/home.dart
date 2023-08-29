import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class Home extends StatelessWidget {
  const Home({Key? key}) : super(key: key);

  static const String routeName = '/';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text('Home', style: TextStyle(color: Colors.white)),
          centerTitle: true,
          backgroundColor: Colors.purple,
        ),
        backgroundColor: Colors.purple,
        body: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Center(
              child: Column(
                children: [
                  const SizedBox(height: 20),
                  ElevatedButton(
                      onPressed: () {
                        context.go('/vehicle');
                      },
                      child: const Text('Vehicle')),
                  const SizedBox(height: 20),
                ],
              ),
            ),
          ),
        ));
  }
}
