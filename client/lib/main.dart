import 'package:flutter/material.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.grey),
        useMaterial3: true,
        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.black,
          foregroundColor: Colors.amberAccent,
        )
      ),
      home: Scaffold(
        appBar: AppBar(
            title: const Text('Transport Predictor'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ElevatedButton(onPressed: () {}, child: const Text("Boton Elevado")),
              SizedBox(height: 20),
              TextButton(onPressed: () {}, child: const Text("Boton de texto")),
              SizedBox(height: 20),
              OutlinedButton(onPressed: () {}, child: const Text("Boton con borde presionado")),
              SizedBox(height: 20),
              IconButton(onPressed: () {}, icon: const Icon(Icons.car_repair)),
              SizedBox(height: 20),
              FloatingActionButton(onPressed: () {}, backgroundColor: Colors.red,foregroundColor: Colors.amber, child: const Icon(Icons.add))
            ],
          ),
        ),
      ),
    );
  }
}
