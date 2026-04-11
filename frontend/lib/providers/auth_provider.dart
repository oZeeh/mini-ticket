import 'package:flutter/material.dart';
import '../services/auth_service.dart';

class AuthProvider extends ChangeNotifier {
  final AuthService _authService = AuthService();

  String? role;
  String? userId;
  bool isLoading = false;
  String? error;

  Future<void> login(String email, String password) async {
    isLoading = true;
    error = null;
    notifyListeners();

    try {
      final payload = await _authService.login(email, password);
      role = payload['role'];
      userId = payload['user_id'];
    } catch (e) {
      error = 'Email ou senha inválidos';
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> logout() async {
    await _authService.logout();
    role = null;
    userId = null;
    notifyListeners();
  }
}
