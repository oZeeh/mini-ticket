import 'dart:convert';
import 'package:frontend/config/config.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../models/ticket.dart';

class TicketService {
  Future<Map<String, String>> _headers() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('token') ?? '';
    return {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $token',
    };
  }

  Future<List<Ticket>> getMyTickets(String userId) async {
    final response = await http.get(
      Uri.parse('${Env.baseUrl}/tickets/user'),
      headers: await _headers(),
    );
    if (response.statusCode == 200) {
      return (jsonDecode(response.body) as List)
          .map((e) => Ticket.fromJson(e))
          .toList();
    }
    throw Exception('Erro ao buscar tickets');
  }

  Future<List<Ticket>> getAllTickets() async {
    final response = await http.get(
      Uri.parse('${Env.baseUrl}/tickets'),
      headers: await _headers(),
    );
    if (response.statusCode == 200) {
      return (jsonDecode(response.body) as List)
          .map((e) => Ticket.fromJson(e))
          .toList();
    }
    throw Exception('Erro ao buscar tickets');
  }

  Future<List<Ticket>> getAssignedTickets() async {
    final response = await http.get(
      Uri.parse('${Env.baseUrl}/tickets/technician'),
      headers: await _headers(),
    );
    if (response.statusCode == 200) {
      return (jsonDecode(response.body) as List)
          .map((e) => Ticket.fromJson(e))
          .toList();
    }
    throw Exception('Erro ao buscar tickets');
  }

  Future<void> createTicket(String title, String description) async {
    final response = await http.post(
      Uri.parse('${Env.baseUrl}/tickets'),
      headers: await _headers(),
      body: jsonEncode({'title': title, 'description': description}),
    );
    if (response.statusCode != 201) throw Exception('Erro ao criar ticket');
  }

  Future<void> assignTicket(String ticketId) async {
    final response = await http.post(
      Uri.parse('${Env.baseUrl}/tickets/$ticketId/assign'),
      headers: await _headers(),
    );
    if (response.statusCode != 200)
      throw Exception('Erro ao se atribuir ao ticket');
  }

  Future<void> updateTicketStatus(String ticketId, String status) async {
    final response = await http.put(
      Uri.parse('${Env.baseUrl}/tickets/$ticketId'),
      headers: await _headers(),
      body: jsonEncode({'status': status}),
    );
    if (response.statusCode != 200) throw Exception('Erro ao atualizar ticket');
  }

  Future<void> deleteTicket(String ticketId) async {
    final response = await http.delete(
      Uri.parse('${Env.baseUrl}/tickets/$ticketId'),
      headers: await _headers(),
    );
    if (response.statusCode != 200) throw Exception('Erro ao deletar ticket');
  }
}
