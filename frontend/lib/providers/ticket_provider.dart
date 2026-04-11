import 'package:flutter/material.dart';
import '../models/ticket.dart';
import '../services/ticket_service.dart';

class TicketProvider extends ChangeNotifier {
  final TicketService _service = TicketService();

  List<Ticket> tickets = [];
  bool isLoading = false;
  String? error;

  Future<void> loadMyTickets(String userId) async {
    _setLoading(true);
    try {
      tickets = await _service.getMyTickets(userId);
    } catch (e) {
      error = e.toString();
    } finally {
      _setLoading(false);
    }
  }

  Future<void> loadAllTickets() async {
    _setLoading(true);
    try {
      tickets = await _service.getAllTickets();
    } catch (e) {
      error = e.toString();
    } finally {
      _setLoading(false);
    }
  }

  Future<void> loadAssignedTickets() async {
    _setLoading(true);
    try {
      tickets = await _service.getAssignedTickets();
    } catch (e) {
      error = e.toString();
    } finally {
      _setLoading(false);
    }
  }

  Future<void> createTicket(String title, String description) async {
    await _service.createTicket(title, description);
  }

  Future<void> assignTicket(String ticketId) async {
    await _service.assignTicket(ticketId);
    notifyListeners();
  }

  Future<void> updateStatus(String ticketId, String status) async {
    await _service.updateTicketStatus(ticketId, status);
    notifyListeners();
  }

  Future<void> deleteTicket(String ticketId) async {
    await _service.deleteTicket(ticketId);
    tickets.removeWhere((t) => t.id == ticketId);
    notifyListeners();
  }

  void _setLoading(bool value) {
    isLoading = value;
    notifyListeners();
  }
}
